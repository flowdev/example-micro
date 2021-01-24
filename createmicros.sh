#!/bin/bash

function createService() {
	local name="$1"
	mkdir "$name"
	cd "$name"
	go mod init
	cp -r ../template/* .
  mv "cmd/main" "cmd/$name"
	find . -name '*.go' -exec sed -i "s/template/$name/g" '{}' \;
  cd ..
}

services="user email catalogue stock creditcard paypal banktransfer buyhistory externalacts internalacts customerstat productstat vouchers shoppingcart"

for srv in $services ; do
	echo "Handling: $srv"
	createService "$srv-cmd"
	createService "$srv-qry"
done

echo "done!"
