module github.com/lucasandre23/microservices-proto/golang/payment

go 1.25.1

require github.com/lucasandre23/microservices-proto/golang/order v0.0.0

replace github.com/lucasandre23/microservices-proto/golang/order => ../order

require github.com/lucasandre23/microservices-proto/golang/order v0.0.0