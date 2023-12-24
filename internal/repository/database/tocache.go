package database

import (
	"github.com/breeeaaad/WB-L0/internal/models"
)

func (d *Database) Alldata() error {
	var orders []models.Order
	rows, err := d.conn.Query(d.context, "select orders.order_uid,orders.track_number,entry,locale,internal_signature,customer_id,delivery_service,shardkey,sm_id,date_created,oof_shard,delivery.name,phone,zip,city,address,region,email,transaction_id,request_id,currency,provider,amount,payment_dt,bank,delivery_cost,goods_total,custom_fee from orders join delivery ON delivery.id = orders.delivery_id join payment ON payment.id = orders.payment_id")
	if err != nil {
		return err
	}
	for rows.Next() {
		var o models.Order
		err = rows.Scan(&o.OrderUid, &o.TrackNumber, &o.Entry, &o.Locale, &o.InternalSignature, &o.CustomerId, &o.DeliveryService, &o.ShardKey, &o.SmId, &o.DateCreated, &o.OofShard, &o.Delivery.Name, &o.Delivery.Phone, &o.Delivery.Zip, &o.Delivery.City, &o.Delivery.Address, &o.Delivery.Region, &o.Delivery.Email, &o.Payment.Transaction, &o.Payment.RequestId, &o.Payment.Currency, &o.Payment.Provider, &o.Payment.Amount, &o.Payment.PaymentDt, &o.Payment.Bank, &o.Payment.DeliveryCost, &o.Payment.GoodsTotal, &o.Payment.CustomFee)
		if err != nil {
			return err
		}
		orders = append(orders, o)
	}
	rows.Close()
	if rows.Err() != nil {
		return err
	}
	for i := 0; i < len(orders); i++ {
		items, err := d.conn.Query(d.context, "select name,sale,size,total_price,nm_id,brand,status,chrt_id,item.track_number,price,rid from item where order_uid=$1", orders[i].OrderUid)
		if err != nil {
			return err
		}
		for items.Next() {
			var item models.Item
			err = items.Scan(&item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmId, &item.Brand, &item.Status, &item.ChrtId, &item.TrackNumber, &item.Price, &item.Rid)
			if err != nil {
				return err
			}
			orders[i].Items = append(orders[i].Items, item)
		}
		items.Close()
		if items.Err() != nil {
			return err
		}
		d.C.Set(orders[i])
	}
	return nil
}
