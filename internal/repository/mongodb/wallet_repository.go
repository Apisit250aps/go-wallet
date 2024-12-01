package mongodb

import (
    "context"
    "time"

    "go-wallet/internal/domain"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type walletRepository struct {
    collection *mongo.Collection
}

func NewWalletRepository(db *mongo.Database) domain.WalletRepository {
    return &walletRepository{
        collection: db.Collection("wallets"),
    }
}

func (r *walletRepository) Create(wallet *domain.Wallet) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    wallet.CreatedAt = time.Now()
    wallet.UpdatedAt = time.Now()

    result, err := r.collection.InsertOne(ctx, wallet)
    if err != nil {
        return err
    }

    wallet.ID = result.InsertedID.(primitive.ObjectID).Hex()
    return nil
}

func (r *walletRepository) GetByID(id string) (*domain.Wallet, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var wallet domain.Wallet
    err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&wallet)
    if err != nil {
        return nil, err
    }

    return &wallet, nil
}

func (r *walletRepository) GetByUserID(userID string) ([]*domain.Wallet, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var wallets []*domain.Wallet
    cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    err = cursor.All(ctx, &wallets)
    if err != nil {
        return nil, err
    }

    return wallets, nil
}

func (r *walletRepository) Update(wallet *domain.Wallet) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    wallet.UpdatedAt = time.Now()

    objectID, err := primitive.ObjectIDFromHex(wallet.ID)
    if err != nil {
        return err
    }

    _, err = r.collection.UpdateOne(
        ctx,
        bson.M{"_id": objectID},
        bson.M{"$set": wallet},
    )
    return err
}

func (r *walletRepository) Delete(id string) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    _, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
    return err
}