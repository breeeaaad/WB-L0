package database

import (
	"github.com/breeeaaad/WB-L0/internal/models"
)

func (d *Database) Create(o models.Order) error {
	tx, err := d.conn.Begin(d.context)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(d.context)
		} else {
			tx.Commit(d.context)
		}
	}()
	var payment_id, delivery_id int
	err = tx.QueryRow(d.context, "insert into delivery(name,phone,zip,city,address,region,email) values($1,$2,$3,$4,$5,$6,$7) returning id", o.Delivery.Name, o.Delivery.Phone, o.Delivery.Zip, o.Delivery.City, o.Delivery.Address, o.Delivery.Region, o.Delivery.Email).Scan(&delivery_id)
	if err != nil {
		return err
	}
	err = tx.QueryRow(d.context, "insert into payment(transaction_id,request_id,currency,provider,amount,payment_dt,bank,delivery_cost,goods_total,custom_fee) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) returning id", o.Payment.Transaction, o.Payment.RequestId, o.Payment.Currency, o.Payment.Provider, o.Payment.Amount, o.Payment.PaymentDt, o.Payment.Bank, o.Payment.DeliveryCost, o.Payment.GoodsTotal, o.Payment.CustomFee).Scan(&payment_id)
	if err != nil {
		return err
	}
	_, err = tx.Exec(d.context, "insert into orders(order_uid,track_number,entry,locale,internal_signature,customer_id,delivery_service,shardkey,sm_id,date_created,oof_shard,delivery_id,payment_id) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)", o.OrderUid, o.TrackNumber, o.Entry, o.Locale, o.InternalSignature, o.CustomerId, o.DeliveryService, o.ShardKey, o.SmId, o.DateCreated, o.OofShard, delivery_id, payment_id)
	if err != nil {
		return err
	}
	for i := 0; i < len(o.Items); i++ {
		_, err = tx.Exec(d.context, "insert into item(name,sale,size,total_price,nm_id,brand,status,chrt_id,track_number,price,rid,order_uid) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)", o.Items[i].Name, o.Items[i].Sale, o.Items[i].Size, o.Items[i].TotalPrice, o.Items[i].NmId, o.Items[i].Brand, o.Items[i].Status, o.Items[i].ChrtId, o.Items[i].TrackNumber, o.Items[i].Price, o.Items[i].Rid, o.OrderUid)
		if err != nil {
			return err
		}
	}
	return nil
}
