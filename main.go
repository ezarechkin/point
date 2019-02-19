package main

import "fmt"
import "log"
import "flag"
import "bufio"
import "net"
import "io"

func main(){

	// Parse launch flags
	flag.Parse()
	
	// Initialize global variables with parsed values
	initializeVariables()
	
	// Run tcp listen loop 
	listenSocket()
}


func initializeVariables(){

	// Build database connection string
	// Point use Postgres database backend with native driver
	DB_CONNECTION_STRING = fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", *DB_USER, *DB_PASS, *DB_HOST, *DB_PORT, *DB_NAME)
}


func listenSocket(){

	var listener	net.Listener
	var err			error

	// Try bind adress for TCP socket
	if listener, err = net.Listen("tcp", *BIND_ADDRESS); err == nil {
		
		log.Printf("NEW Listener: address: %s", *BIND_ADDRESS)

		// Prepare instance auto destruction after use
		defer func(){

			log.Printf("CLOSE Listener: address: %s", *BIND_ADDRESS)
			listener.Close()
		}()

		// Wait for connections
		for {

			var connection	net.Conn		// Define new connection variable for every new connection
			var err			error			// Define new error variable for every new connection
			
			// Accept connection from listener
			if connection, err = listener.Accept(); err == nil {
				
				log.Printf("NEW Connection: address: %s", connection.RemoteAddr())

				// We dont need threads buffer for comunications now
				// So just run goroutin
				go handleConnection(connection)
			} else {
				
				// Log problem connection without exit
				log.Println(err)
				break
			}
		}
	} else {
		
		// Exit if can't assing address :(
		log.Fatal(err)
	}
}

func handleConnection(connection net.Conn){

	// Prepare instance auto destruction after use
	defer func(){
		log.Printf("CLOSE Connection: address: %s", connection.RemoteAddr())
		connection.Close()
	}()

	// Wait for messages
	for {
			
		var message	string
		var err		error

		// Accept message from connection
		if message, err = bufio.NewReader(connection).ReadString('\n'); err == nil {
			
			var output	string
			var err		error

			log.Printf("NEW Query: address: %s", connection.RemoteAddr())

			// Process message
			if output, err = authenticateUser(&message); err == nil {
						
				io.WriteString(connection, fmt.Sprintf("%s\n", output))						
			} else {
						
				log.Println(err)
				io.WriteString(connection, fmt.Sprintf("%s\n", output))
			}
		} else {

			// Close connection and stop buf reader
			// if error or client disconnect
		 	log.Println(err)
			break
		}
	}
}