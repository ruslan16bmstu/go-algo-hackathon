package rpc_server

import "pocket-trader-backend/db"

type RPCServer struct {
	// proto interface
	DB db.Api // interface for database
}
