package tools

import (
	"log"
	"time"

	"github.com/mattpaletta/AbilitySoftwareGroup468/networks"

	"github.com/mattpaletta/AbilitySoftwareGroup468/common"
)

type TriggerManager struct {
	c      Cache
	db     *CacheDB
	logger networks.Logger
}

func (tm *TriggerManager) Start() {
	go func() {
		for {
			log.Println("Executing triggers...")
			trigs, err := tm.db.Triggers.GetAll()
			if err == nil {
				txns := make([]*common.PendingTxn, 0)
				for _, trig := range trigs {
					txn := tm.processTrigger(trig)
					if txn != nil {
						txns = append(txns, txn)
					}
				}
				if len(txns) > 0 {
					log.Printf("Resolving %d transactions\n", len(txns))
					err = tm.db.Users.BulkTransaction(txns, false)
					if err != nil {
						log.Println(err)
					}
					err = tm.db.Transactions.BulkLog(txns, true)
					if err != nil {
						log.Println(err)
					}
					err = tm.db.Triggers.BulkClose(txns)
					if err != nil {
						log.Println(err)
					}
				}
			}
			time.Sleep(time.Minute)
		}
	}()
}

func (tm *TriggerManager) processTrigger(t common.Trigger) *common.PendingTxn {
	quote, err := tm.c.GetQuote(t.Stock, t.UserId, t.TransactionID)
	if err != nil {
		return nil
	}
	isTriggered := (t.Type == "BUY" && quote.Quote <= t.When) ||
		(t.Type == "SELL" && quote.Quote >= t.When)
	if !isTriggered {
		return nil
	}
	commandType := common.SET_BUY_TRIGGER
	action := "remove"
	if t.Type == "SELL" {
		commandType = common.SET_SELL_TRIGGER
		action = "add"
	}
	go tm.logger.AccountTransaction(t.UserId, t.Amount, action, t.TransactionID)
	go tm.logger.SystemEvent(&common.Command{
		C_type:        commandType,
		UserId:        t.UserId,
		StockSymbol:   t.Stock,
		Amount:        t.Amount,
		TransactionID: t.TransactionID})

	shares := t.Shares
	if t.Type == "BUY" {
		shares = int(t.Amount / quote.Quote)
	}
	price := int64(shares) * quote.Quote
	return &common.PendingTxn{
		UserId:   t.UserId,
		Price:    price,
		Reserved: t.Amount,
		Shares:   shares,
		Type:     t.Type,
		Stock:    t.Stock,
	}
}

func NewTrigMan(c Cache, db *CacheDB, l networks.Logger) *TriggerManager {
	return &TriggerManager{c, db, l}
}
