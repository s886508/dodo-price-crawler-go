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
    "Detail": "南港展覽館2站2月份活動展覽期間，臨停費率調整如下：2/21~2/26        工具機展TIMTOS x TMTS2022：$60/H,無當日當次最高上限"
  },
  {
    "Station": "南港展覽館站",
    "Detail": "南港展覽館站2月份活動展覽期間，臨停費率調整如下：2/21~2/26        工具機展TIMTOS x TMTS2022：$60/H,無當日當次最高上限"
  }
]
```
