package mongodb

import (
    "context"
    "time"

    "go-wallet/internal/domain"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
    collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) domain.UserRepository {
    return &userRepository{
        collection: db.Collection("users"),
    }
}

func (r *userRepository) Create(user *domain.User) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()

    result, err := r.collection.InsertOne(ctx, user)
    if err != nil {
        return err
    }

    user.ID = result.InsertedID.(primitive.ObjectID).Hex()
    return nil
}

func (r *userRepository) GetByID(id string) (*domain.User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var user domain.User
    err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (r *userRepository) GetByUsername(username string) (*domain.User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var user domain.User
    err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (r *userRepository) Update(user *domain.User) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    user.UpdatedAt = time.Now()

    objectID, err := primitive.ObjectIDFromHex(user.ID)
    if err != nil {
        return err
    }

    _, err = r.collection.UpdateOne(
        ctx,
        bson.M{"_id": objectID},
        bson.M{"$set": user},
    )
    return err
}