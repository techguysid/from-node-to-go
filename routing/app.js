let http = require('http');
http.createServer(function (req, res) {
 res.setHeader('Content-Type', 'application/json');
 let url = req.url;
 if(url ==='/a'){
    res.write(JSON.stringify({'message':'You have hit /a'}))
    res.end()
 }else{
    res.write(JSON.stringify({'message':'You have hit something other than /a'}))
    res.end()
 }}).listen(8081, function(){
 console.log("server start at port 8081"); //the server object listens on port 3000
});