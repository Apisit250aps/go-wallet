package domain

type UserUsecase interface {
    Register(user *User) error
    Login(username, password string) (string, error) // returns JWT token
    GetUserByID(id string) (*User, error)
}

type WalletUsecase interface {
    CreateTransaction(wallet *Wallet) error
    GetTransactions(userID string) ([]*Wallet, error)
    GetTransaction(id string) (*Wallet, error)
    UpdateTransaction(wallet *Wallet) error
    DeleteTransaction(id string) error
}