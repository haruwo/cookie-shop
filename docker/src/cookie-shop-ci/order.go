package main

import (
	"cookie-shop-ci/lib/files"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewOrder() *cobra.Command {
	order := &cobra.Command{
		Use:   "order",
		Short: "Order commands",
	}

	order.AddCommand(newConfirmOrder())

	return order
}

func newConfirmOrder() *cobra.Command {
	return &cobra.Command{
		Use:   "confirm orderID [targetDir]",
		Short: "Confirm order",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("must presence orderID")
			}

			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			orderID := args[0]
			targetDir := "."
			if len(args) >= 2 {
				targetDir = args[1]
			}

			fs, err := files.Open(targetDir)
			if err != nil {
				return err
			}

			order := fs.FindOrder(orderID)
			if order == nil {
				return fmt.Errorf("can't find order in orders. orderID:%s", orderID)
			}

			var billing OrderBilling
			var items int64
			for _, i := range order.Items {
				item := fs.FindItem(i.ID)
				if item == nil {
					return fmt.Errorf("can't find item in items. itemID:%s", i.ID)
				}
				billing.Products += item.Price * i.Amount
				items += i.Amount
			}

			billing.Shipping = ((items + 5) / 6) * 198

			total := billing.Products + billing.Shipping
			billing.Tax = total / 10
			billing.Total = total + billing.Tax

			co := OrderInfo{Billing: &billing}

			w := json.NewEncoder(os.Stdout)
			if err := w.Encode(co); err != nil {
				return err
			}

			return nil
		},
	}
}

type OrderInfo struct {
	Billing *OrderBilling `json:"billing"`
}

type OrderBilling struct {
	Products int64 `json:"products"`
	Shipping int64 `json:"shipping"`
	Tax      int64 `json:"tax"`
	Total    int64 `json:"total"`
}
