package usecase

import (
    "errors"
    "go-wallet/internal/domain"
)

type walletUsecase struct {
    walletRepo domain.WalletRepository
}

func NewWalletUsecase(wr domain.WalletRepository) domain.WalletUsecase {
    return &walletUsecase{
        walletRepo: wr,
    }
}

func (w *walletUsecase) CreateTransaction(wallet *domain.Wallet) error {
    if wallet.Amount <= 0 {
        return errors.New("amount must be greater than 0")
    }

    if wallet.Type != "income" && wallet.Type != "expense" {
        return errors.New("invalid transaction type")
    }

    return w.walletRepo.Create(wallet)
}

func (w *walletUsecase) GetTransactions(userID string) ([]*domain.Wallet, error) {
    return w.walletRepo.GetByUserID(userID)
}

func (w *walletUsecase) GetTransaction(id string) (*domain.Wallet, error) {
    return w.walletRepo.GetByID(id)
}

func (w *walletUsecase) UpdateTransaction(wallet *domain.Wallet) error {
    if wallet.Amount <= 0 {
        return errors.New("amount must be greater than 0")
    }

    if wallet.Type != "income" && wallet.Type != "expense" {
        return errors.New("invalid transaction type")
    }

    return w.walletRepo.Update(wallet)
}

func (w *walletUsecase) DeleteTransaction(id string) error {
    return w.walletRepo.Delete(id)
}