# snippetbox

《Let’s Go: Learn to build professional web applications with Go》


## Chapter 2

curl -i -X POST http://localhost:4000/snippet/create
curl -i -X PUT http://localhost:4000/snippet/create

http://localhost:4000/snippet?id=1


mkdir -p cmd/web pkg ui/html ui/static

go run ./cmd/web

curl https://www.alexedwards.net/static/sb.v130.tar.gz | tar -xvz -C ./ui/static

> Range requests are fully supported. This is great if your application is servinglarge files and you want to support resumable downloads. You can see thisfunctionality in action if you use curl to request bytes 100-199 of the logo.png file,like so:

curl -i -H "Range: bytes=100-199" --output - http://localhost:4000/static/img/logo.png

## Chapter 3 Managing Configuration Settings

