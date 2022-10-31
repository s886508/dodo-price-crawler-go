# Description
Retrieve dodo home(嘟嘟房) special prices from the official site

# Build
```go
go build -o <output_file> .\main.go
```

# Add interested parking stations
Add station names to `configs/config.json`.
```json
{
  "url": "https://www.dodohome.com.tw/p1_news.aspx",
  "stations": [
    "南港展覽館站",
    "南港展覽館2站"
  ]
}
```

# Results in JSON
```json
[
  {
    "Station": "南港展覽館2站",
    "Detail": [
      "南港展覽館2站10月份活動展覽期間，臨停費率調整為：$60/H，當日當次最高上限$350~",
      "9/27~10/1     台北國際塑橡膠工業展",
      "10/12~10/16   2022台北紡織展/2022展昭世界貓咪博覽會"
    ]
  },
  {
    "Station": "南港展覽館站",
    "Detail": [
      "南港展覽館站10月份活動展覽期間，臨停費率調整如下~",
      "9/27~10/1     台北國際塑橡膠工業展：$60/H，無最高上限",
      "10/8          瑞納2022全球菁英表揚大會：$60/H，當日當次最高上限$350"
    ]
  }
]
```
