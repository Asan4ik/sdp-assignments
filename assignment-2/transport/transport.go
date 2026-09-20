package transport

/* Transport is the product interface. Every delivery method has to
   know how to deliver something and describe what it did. */
type Transport interface {
	Deliver() string
}
