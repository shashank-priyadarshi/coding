Facade is a structural design pattern that provides a simplified (but limited) interface to a complex system of classes, library or framework.

While Facade decreases the overall complexity of the application, it also helps to move unwanted dependencies to one place.

Conceptual Example
It’s easy to underestimate the complexities that happen behind the scenes when you order a pizza using your credit card. There are dozens of subsystems that are acting in this process. Here’s just a shortlist of them:

-Check account 
-Check security PIN
-Credit/debit balance
-Make ledger entry
-Send notification
In a complex system like this, it’s easy to get lost and easy to break stuff if you’re doing something wrong. That’s why there’s a concept of the Facade pattern: a thing that lets the client work with dozens of components using a simple interface. The client only needs to enter the card details, the security pin, the amount to pay, and the operation type. The Facade directs further communications with various components without exposing the client to internal complexities.

### walletFacade.go: Facade
```go
package main

import "fmt"

type WalletFacade struct {
    account      *Account
    wallet       *Wallet
    securityCode *SecurityCode
    notification *Notification
    ledger       *Ledger
}

func newWalletFacade(accountID string, code int) *WalletFacade {
    fmt.Println("Starting create account")
    walletFacacde := &WalletFacade{
        account:      newAccount(accountID),
        securityCode: newSecurityCode(code),
        wallet:       newWallet(),
        notification: &Notification{},
        ledger:       &Ledger{},
    }
    fmt.Println("Account created")
    return walletFacacde
}

func (w *WalletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
    fmt.Println("Starting add money to wallet")
    err := w.account.checkAccount(accountID)
    if err != nil {
        return err
    }
    err = w.securityCode.checkCode(securityCode)
    if err != nil {
        return err
    }
    w.wallet.creditBalance(amount)
    w.notification.sendWalletCreditNotification()
    w.ledger.makeEntry(accountID, "credit", amount)
    return nil
}

func (w *WalletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
    fmt.Println("Starting debit money from wallet")
    err := w.account.checkAccount(accountID)
    if err != nil {
        return err
    }

    err = w.securityCode.checkCode(securityCode)
    if err != nil {
        return err
    }
    err = w.wallet.debitBalance(amount)
    if err != nil {
        return err
    }
    w.notification.sendWalletDebitNotification()
    w.ledger.makeEntry(accountID, "debit", amount)
    return nil
}
```
### account.go: Complex subsystem parts
```go
package main

import "fmt"

type Account struct {
    name string
}

func newAccount(accountName string) *Account {
    return &Account{
        name: accountName,
    }
}

func (a *Account) checkAccount(accountName string) error {
    if a.name != accountName {
        return fmt.Errorf("Account Name is incorrect")
    }
    fmt.Println("Account Verified")
    return nil
}
```
### securityCode.go: Complex subsystem parts
```go
package main

import "fmt"

type SecurityCode struct {
    code int
}

func newSecurityCode(code int) *SecurityCode {
    return &SecurityCode{
        code: code,
    }
}

func (s *SecurityCode) checkCode(incomingCode int) error {
    if s.code != incomingCode {
        return fmt.Errorf("Security Code is incorrect")
    }
    fmt.Println("SecurityCode Verified")
    return nil
}
```
### wallet.go: Complex subsystem parts
```go
package main

import "fmt"

type Wallet struct {
    balance int
}

func newWallet() *Wallet {
    return &Wallet{
        balance: 0,
    }
}

func (w *Wallet) creditBalance(amount int) {
    w.balance += amount
    fmt.Println("Wallet balance added successfully")
    return
}

func (w *Wallet) debitBalance(amount int) error {
    if w.balance < amount {
        return fmt.Errorf("Balance is not sufficient")
    }
    fmt.Println("Wallet balance is Sufficient")
    w.balance = w.balance - amount
    return nil
}
```
### ledger.go: Complex subsystem parts
```go
package main

import "fmt"

type Ledger struct {
}

func (s *Ledger) makeEntry(accountID, txnType string, amount int) {
    fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
    return
}
```
### notification.go: Complex subsystem parts
```go
package main

import "fmt"

type Notification struct {
}

func (n *Notification) sendWalletCreditNotification() {
    fmt.Println("Sending wallet credit notification")
}

func (n *Notification) sendWalletDebitNotification() {
    fmt.Println("Sending wallet debit notification")
}
```
### main.go: Client code
```go
package main

import (
    "fmt"
    "log"
)

func main() {
    fmt.Println()
    walletFacade := newWalletFacade("abc", 1234)
    fmt.Println()

    err := walletFacade.addMoneyToWallet("abc", 1234, 10)
    if err != nil {
        log.Fatalf("Error: %s\n", err.Error())
    }

    fmt.Println()
    err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
    if err != nil {
        log.Fatalf("Error: %s\n", err.Error())
    }
}
```
### output.txt: Execution result
```
Starting create account
Account created

Starting add money to wallet
Account Verified
SecurityCode Verified
Wallet balance added successfully
Sending wallet credit notification
Make ledger entry for accountId abc with txnType credit for amount 10

Starting debit money from wallet
Account Verified
SecurityCode Verified
Wallet balance is Sufficient
Sending wallet debit notification
Make ledger entry for accountId abc with txnType debit for amount 5
```