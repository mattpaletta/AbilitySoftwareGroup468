package logging

import (
	"encoding/xml"
)

type Args struct {
	Action, Command, Cryptokey, DebugMessage, ErrorMessage, FileName, Server, StockSymbol, Username string
	Funds, Price, TransactionNum                                                                    int64
	Timestamp, QuoteServerTime                                                                      uint64
}

type UserCommand struct {
	XMLName        xml.Name `xml:"userCommand"`
	Timestamp      uint64   `xml:"timestamp"`
	TransactionNum int64    `xml:"transactionNum"`
	Command        string   `xml:"command"`
	Username       string   `xml:"username,omitempty"`
	StockSymbol    string   `xml:"stockSymbol,omitempty"`
	Filename       string   `xml:"filename,omitempty"`
	Funds          int64    `xml:"funds,omitempty"`
}

type QuoteServer struct {
	XMLName         xml.Name `xml:"quoteServer"`
	Timestamp       uint64   `xml:"timestamp"`
	TransactionNum  int64    `xml:"transactionNum"`
	Price           int64    `xml:"price"`
	StockSymbol     string   `xml:"stockSymbol"`
	Username        string   `xml:"username"`
	QuoteServerTime uint64   `xml:"quoteServerTime"`
	Cryptokey       string   `xml:"cryptokey"`
}

type AccountTransaction struct {
	XMLName        xml.Name `xml:"accountTransaction"`
	Timestamp      uint64   `xml:"timestamp"`
	TransactionNum int64    `xml:"transactionNum"`
	Action         string   `xml:"action"`
	Username       string   `xml:"username"`
	Funds          int64    `xml:"funds"`
}

type SystemEvent struct {
	XMLName        xml.Name `xml:"systemEvent"`
	Timestamp      uint64   `xml:"timestamp"`
	TransactionNum int64    `xml:"transactionNum"`
	Command        string   `xml:"command"`
	Username       string   `xml:"username, omitempty"`
	StockSymbol    string   `xml:"stockSymbol, omitempty"`
	Filename       string   `xml:"filename, omitempty"`
	Funds          int64    `xml:"funds, omitempty"`
}

type ErrorEvent struct {
	XMLName        xml.Name `xml:"errorEvent"`
	Timestamp      uint64   `xml:"timestamp"`
	TransactionNum int64    `xml:"transactionNum"`
	Command        string   `xml:"command"`
	Username       string   `xml:"username, omitempty"`
	StockSymbol    string   `xml:"stockSymbol, omitempty"`
	Filename       string   `xml:"filename, omitempty"`
	Funds          int64    `xml:"funds, omitempty"`
	ErrorMessage   string   `xml:"errorMessage, omitempty"`
}

type DebugEvent struct {
	XMLName        xml.Name `xml:"debugEvent"`
	Timestamp      uint64   `xml:"timestamp"`
	TransactionNum int64    `xml:"transactionNum"`
	Command        string   `xml:"command"`
	Username       string   `xml:"username, omitempty"`
	StockSymbol    string   `xml:"stockSymbol, omitempty"`
	Filename       string   `xml:"filename, omitempty"`
	Funds          int64    `xml:"funds, omitempty"`
	DebugMessage   string   `xml:"debugMessage, omitempty"`
}
