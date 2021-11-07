package types

//Money - денежная сумма в мин.еденицах (центы, копейки, дирамы и тд)
type Money int64

// Currency - категория в которой был совершен платеж(авто, аптеки, рестораны и тд)
type Category string

//Payment - информация о платеже
type Payment struct {
	ID			int
	Amount 		Money
	Category 	Category
}
