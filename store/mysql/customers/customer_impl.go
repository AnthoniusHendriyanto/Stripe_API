package customers

import (
	"database/sql"
	"fmt"

	"stripe_api/models/request"
	"stripe_api/models/response"
)

func (o *customerClient) InsertCustomers(req request.InsertCustomers) error {
	query, err := o.DB.Prepare(`INSERT INTO customer SET customerId=?, name=?, phone=?, status=?;`)
	if err != nil {
		fmt.Printf("Error preparing customers : %v", err)
		return err
	}

	defer func(query *sql.Stmt) {
		_ = query.Close()
	}(query)

	_, errExec := query.Exec(req.CustomerID, req.Name, req.PhoneNumber, req.Status)
	if errExec != nil {
		fmt.Printf("Error executing customers : %v", errExec)
		return errExec
	}

	return nil
}

func (o *customerClient) GetCustomers(req request.GetCustomer) (
	*response.GetCustomerResponse, error) {

	res := &response.GetCustomerResponse{}

	err := o.DB.QueryRow(`SELECT c.customerId, c.name, c.phone, c.status FROM customer c WHERE c.phone = ?;`,
		req.PhoneNumber).Scan(&res.CustomerID, &res.Name, &res.PhoneNumber, &res.Status)

	if err != nil {
		fmt.Printf("Error getting customers: %v", err)

		return nil, err
	}

	return res, nil
}

func (o *customerClient) InsertCard(req request.CreateCardDatabaseReq) error {
	query, err := o.DB.Prepare(`INSERT INTO card SET cardId=?, brand=?, customerId=?;`)
	if err != nil {
		fmt.Printf("Error preparing customers : %v", err)
		return err
	}

	defer func(query *sql.Stmt) {
		_ = query.Close()
	}(query)

	_, errExec := query.Exec(req.CardID, req.Brand, req.CustomerID)
	if errExec != nil {
		fmt.Printf("Error executing customers : %v", errExec)
		return errExec
	}

	return nil
}

func (o *customerClient) GetCardDetails(req request.GetListCardsRequest) (
	*response.GetListCardInternalResponse, error) {

	var res []*response.GetListCards

	queryString := `SELECT c.cardId, c.brand, c2.name, c2.phone, cs.description
        FROM card c INNER JOIN customer c2 ON c.customerId = c2.customerId 
        INNER JOIN customer_status cs ON c2.status = cs.status
		WHERE 1=1`

	if req.Brand != "" {
		queryString += " AND c.brand = '" + req.Brand + "'"
	}
	if req.PhoneNumber != "" {
		queryString += " AND c2.phone = '" + req.PhoneNumber + "'"
	}
	if req.Brand != "" && req.PhoneNumber != "" {
		queryString += " AND c.brand = '" + req.Brand + "' AND c2.phone = '" + req.PhoneNumber + "'"
	}

	fmt.Println("Query : " + queryString)

	query, err := o.DB.Query(queryString)

	if err != nil {
		fmt.Printf("Error getting customers: %v", err)

		return nil, err
	}

	for query.Next() {
		infoCard := &response.GetListCards{}

		err := query.Scan(
			&infoCard.CardID, &infoCard.BrandCard, &infoCard.CustomerName, &infoCard.PhoneNumber, &infoCard.Status,
		)
		if err != nil {
			fmt.Printf("Error while scanning query: %v", err)
			return nil, err
		}

		res = append(res, infoCard)
	}

	responseList := &response.GetListCardInternalResponse{
		ListCard: res,
	}

	return responseList, nil
}
