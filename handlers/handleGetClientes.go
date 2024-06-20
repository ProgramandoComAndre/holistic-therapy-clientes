package handlers

import(
	"net/http"
	"context"
	"fmt"
	"clientes/repositories"
	"github.com/jackc/pgx/v5"
	"os"
	"encoding/json"
	"clientes/middleware"
)

type ErrorType struct {
	Message string
}

func HandleGetClientes(w http.ResponseWriter, r *http.Request) {
	userClaims, success := middleware.VerifyAuth(r)
	if(!success) {
		err := ErrorType {

		}
		err.Message = "Invalid Token"
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err)
	    return 
	}
	fmt.Printf("%s %d", (*userClaims).Username, (*userClaims).RoleId)

	roleId := (*userClaims).RoleId
	if(roleId == 1) {
		//administrator
		err := ErrorType {

		}
		err.Message = "Protected data. Administrator does not have access to customers"
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
	    return
	}

	conn, err := pgx.Connect(context.Background(), "postgres://clientes:clientes@localhost:5433/clientes")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer conn.Close(context.Background())

	clientes, err := repositories.GetClientes(conn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(clientes)

}