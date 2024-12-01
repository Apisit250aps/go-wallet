package domain

type UserRepository interface {
    Create(user *User) error
    GetByID(id string) (*User, error)
    GetByUsername(username string) (*User, error)
    Update(user *User) error
}

type WalletRepository interface {
    Create(wallet *Wallet) error
    GetByID(id string) (*Wallet, error)
    GetByUserID(userID string) ([]*Wallet, error)
    Update(wallet *Wallet) error
    Delete(id string) error
}