package procon_data

import(
	"fmt"
	"github.com/gorilla/websocket"
)



/* Websocket Message Data */
type Msg struct {
	Jwt string `json:"jwt"`
	Type string `json:"type"`
	Data string	`json:"data"`
}

func SendMsg(j string, t string, d string, c *websocket.Conn) {
	m := Msg{j, t, d};
	if err := c.WriteJSON(m); err != nil {
		fmt.Println(err)
	}

	//mm, _ := json.Marshal(m);
	//fmt.Println(string(mm));
}

/* Websocket Pool Objects */
type FatClient struct {
    Id  string
    Conn *websocket.Conn
    Pool *Pool
}

type Pool struct {
    Register   chan *FatClient
    Unregister chan *FatClient
    Clients    map[*FatClient]bool
}

func NewPool() *Pool {
    return &Pool{
        Register:   make(chan *FatClient),
        Unregister: make(chan *FatClient),
        Clients:    make(map[*FatClient]bool),
    }
}

func (pool *Pool) Start() {
	fmt.Println("Websocket Pool Starting...")
	
	for {
        select {
	        case client := <-pool.Register:
	        	pool.Clients[client] = true
	        	fmt.Println("Size of Connection Pool: ", len(pool.Clients))	        	
	        	break;
		    case client := <-pool.Unregister:
		    	delete(pool.Clients, client)
		    	fmt.Println("Size of Connection Pool: ", len(pool.Clients))
		    	break;
	        default:
	        	break;  
	    }   
	}		
}
