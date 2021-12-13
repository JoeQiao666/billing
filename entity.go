package billing

import (
	"fmt"
	"log"

	domain "github.com/JoeQiao666/billing/persistence"
	"github.com/cloudstateio/go-support/cloudstate/eventsourced"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

type BillingCenter struct {
	prices *domain.Prices
}

func NewBillingCenter(eventsourced.EntityID) eventsourced.EntityHandler {
	return &BillingCenter{
		prices: new(domain.Prices),
	}
}

func (b *BillingCenter) HandleEvent(ctx *eventsourced.Context, event interface{}) error {
	switch e := event.(type) {
	case *domain.PricesAddedOrUpdated:
		return b.PricesAddedOrUpdated(e)
	case *domain.PricesDeleted:
		return b.PricesDeleted(e)
	default:
		return nil
	}
}

func (b *BillingCenter) HandleCommand(ctx *eventsourced.Context, name string, cmd proto.Message) (proto.Message, error) {
	switch c := cmd.(type) {
	case *AddOrUpdatePrices:
		return b.AddOrUpdatePrices(ctx, c)
	case *DeletePrices:
		return b.DeletePrices(ctx, c)
	case *QueryPrices:
		return b.GetPrices(ctx, c)
	case *CalculatePrice:
		return b.CalculatePrice(ctx, c)
	default:
		return nil, nil
	}
}

func (b *BillingCenter) GetPrices(*eventsourced.Context, *QueryPrices) (*Prices, error) {
	prices := []*Price{}
	for _, v := range b.prices.Prices {
		prices = append(prices, &Price{
			ProductId:   v.GetProductId(),
			ProductName: v.GetProductName(),
			Price:       v.GetPrice(),
		})
	}
	return &Prices{
		Prices: prices,
	}, nil
}

func (b *BillingCenter) getPriceById(id string) (int32, error) {
	for _, v := range b.prices.GetPrices() {
		if v.GetProductId() == id {
			return v.GetPrice(), nil
		}
	}
	return 0, fmt.Errorf("there is no price info of %s", id)
}

func (b *BillingCenter) CalculatePrice(c *eventsourced.Context, items *CalculatePrice) (*Resp, error) {
	var totalPrice int32 = 0
	for _, v := range items.GetProducts() {
		price, err := b.getPriceById(v.GetProductId())
		if err != nil {
			log.Println(err.Error())
			return &Resp{}, err
		}
		totalPrice = totalPrice + price*v.GetQuantity()
	}
	return &Resp{
		ProductsWithTotalPrice: &ProductsWithTotalPrice{
			TotalPrice: totalPrice,
			Products:   items.GetProducts(),
		},
	}, nil
}

func (b *BillingCenter) addOrUpdate(price *domain.Price) {
	for _, v := range b.prices.Prices {
		if v.GetProductId() == price.GetProductId() {
			v.Price = price.GetPrice()
			return
		}
	}
	b.prices.Prices = append(b.prices.Prices, price)
}

func (b *BillingCenter) AddOrUpdatePrices(ctx *eventsourced.Context, prices *AddOrUpdatePrices) (*empty.Empty, error) {
	items := prices.GetPrices()
	addedItems := []*domain.Price{}
	for _, v := range items {
		price := v.GetPrice()
		if price < 0 {
			return nil, fmt.Errorf("can't add negative price %d of commodity %s", price, v.GetProductName())
		}
		addedItems = append(addedItems, &domain.Price{
			ProductId:   v.GetProductId(),
			ProductName: v.GetProductName(),
			Price:       price,
		})
	}
	ctx.Emit(&domain.PricesAddedOrUpdated{
		Prices: addedItems,
	})
	return &empty.Empty{}, nil
}

func (b *BillingCenter) hasItem(price *Price) error {
	for _, v := range b.prices.Prices {
		if v.GetProductId() == price.GetProductId() {
			return nil
		}
	}
	return fmt.Errorf("there is no price info of %s", price.GetProductName())
}

func (b *BillingCenter) DeletePrices(ctx *eventsourced.Context, prices *DeletePrices) (*empty.Empty, error) {
	items := prices.GetPrices()
	deletedItems := []*domain.Price{}
	for _, v := range items {
		err := b.hasItem(v)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		deletedItems = append(deletedItems, &domain.Price{
			ProductId:   v.GetProductId(),
			ProductName: v.GetProductName(),
			Price:       v.GetPrice(),
		})
	}
	ctx.Emit(&domain.PricesDeleted{
		Prices: deletedItems,
	})
	return &empty.Empty{}, nil
}

func (b *BillingCenter) PricesAddedOrUpdated(prices *domain.PricesAddedOrUpdated) error {
	for _, v := range prices.GetPrices() {
		b.addOrUpdate(v)
	}
	return nil
}

func (b *BillingCenter) PricesDeleted(prices *domain.PricesDeleted) error {
	for _, deleteItem := range prices.GetPrices() {
		for idx, exist := range b.prices.GetPrices() {
			if deleteItem.GetProductId() == exist.GetProductId() {
				b.prices.Prices = append(b.prices.Prices[:idx], b.prices.Prices[idx+1:]...)
			}
		}
	}
	return nil
}
