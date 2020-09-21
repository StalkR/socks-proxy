all:
	go build .
install:
	mkdir -p $(DESTDIR)/usr/bin
	cp socks-proxy $(DESTDIR)/usr/bin
	chmod 755 $(DESTDIR)/usr/bin/socks-proxy
clean:  
	rm -f socks-proxy
