```
curl -X POST http://localhost:9001/post/file/upload/abc123/application-json \
  -F "file=@./aa.txt" \
  -F "sutra_name=论语" \
  -F "item_name=不食嗟来之食" \
  -F "item_number=3" \
  -H "Content-Type: multipart/form-data"
  
```


curl -X POST http://localhost:9001/post/file/upload/abc123/application-json \
  -F "sutra_name=论语&item_number=3&item_name=不食嗟来之食&file=@./aa.txt"\
  -H "Content-Type: multipart/form-data"
  
