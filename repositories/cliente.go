package repositories

import (
	"context"
	"github.com/jackc/pgx/v5"
	"clientes/models" // Import the models package
)

func GetClientes(conn *pgx.Conn)([]models.Cliente, error) {
   rows, err := conn.Query(context.Background(), "SELECT ID, COALESCE(clientes.nome, '') AS nome, COALESCE(clientes.email, '') AS email, COALESCE(clientes.telefone, '') AS telefone, COALESCE(clientes.morada, '') AS morada ,COALESCE(clientes.doencas, '') AS doencas, COALESCE(clientes.outros, '') AS outros FROM clientes")
   if err != nil {
	return nil, err
   }
   defer rows.Close()

   var clientes []models.Cliente
   for rows.Next() {
	var cliente models.Cliente
	if err := rows.Scan(&cliente.ID, &cliente.Nome, &cliente.Email, &cliente.Telefone, &cliente.Morada, &cliente.Doencas, &cliente.Outros); err != nil {
		return clientes, err
	}
	clientes = append(clientes, cliente)
   }

   if err = rows.Err(); err != nil {
		return clientes, err
   }

   return clientes, nil
}