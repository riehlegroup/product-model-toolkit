# run the server
./run-server.sh

# import 
./cli --type=scanner --path=output/phpScanner.json

# export
./cli export --id=1 --path=res.spdx --type=spdx

# diff
./cli diff path --first=1.spdx --second=2.spdx

# scanner
./cli scanner --out=$PWD/output --name=phpscanner --source=$PWD/source/laravel

# search
./cli scanner --out=search.spdx --name=hello-package 
