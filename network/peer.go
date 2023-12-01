package network

import (
	"encoding/json"
	"net/http"
)

type Peer struct {
	Blockchain *Blockchain
}

func NewPeer(addr string, bc *Blockchain) *Peer {
	peer := &Peer{Blockchain: bc}
	http.HandleFunc("/blocks", peer.handleBlocks)
	return peer
}

func (p *Peer) handleBlocks(w http.ResponseWriter, r *http.Request) {
	blocks := p.Blockchain.GetBlocks()
	respondWithJSON(w, http.StatusOK, blocks)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
