 package card
 
 import (
	 "bank/pkg/bank/types"
 )
 func PaymentSources(cards []types.Card) []types.PaymentSource {
	var payment []types.PaymentSource
	 for _, card := range cards {
		 if card.Balance >0 && card.Active{
			 payment = append(payment,types.PaymentSource{Number:string (card.PAN), Balance:card.Balance}) 
		 }
	 }
	 return payment
 }