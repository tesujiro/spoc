# spoc
Spotify web api CLI

# Install
```
> go install github.com/tesujiro/spoc
```

# Setup
```
ClientID=XXXXXXXXXXXXXXXXXXXXXXXXXXXX
ClientSecret=XXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

# Usage Example
```
> spoc
Usage:
        spoc search [-id] album(s)|artist(s)|playlist(s)|track(s) [keyword]
        spoc [-id] get album(s) [album_id]+
        spoc [-id] get feature(s) [track_id]+
        spoc [-id] get profile [user_id]+
        spoc [-id] get playlist [playlist_id]+
        spoc [-id] get playlists [user_id]+
        spoc [-id] get track(s) [track_id]+
        spoc [-id] list device(s)
        spoc [-id] list playlist(s)
        spoc [-id] list profile
        spoc play [device_id]
        spoc pause [device_id]
        spoc playing [device_id]
        spoc seek [position_ms] [device_id]
        spoc play next [device_id]
        spoc play previous [device_id]
> spoc search artist つのだ
Total: 7
Artists[0]:     1nMq5kf0MNoyReLXbHG9ky  name:Kenichi Tsunoda Big Band   
Artists[1]:     5quBwjQ8JGOfyqN1ZO5Vr4  name:石丸幹二&つのだたかし    
Artists[2]:     0EEoCOufYQhyFawRDiyXRd  name:Hiro Tsunoda       
Artists[3]:     6Lr9vTJ50JwDVIsbZOhUT7  name:つのだたかし 
Artists[4]:     5GmtQhIhCw2NFHLs7zV1SL  name:つのだたかし 
Artists[5]:     3LOXKIDsx3TCg1ADlap4q1  name:つのだ☆ひろ&大橋 純子   
Artists[6]:     6dGiZRfecHPOU7n5ZaafHV  name:つのだ健
> spoc search album mattn
Total: 14
Album[0]:       0o9OjYh1Ep8XPArRoQMoXt  name:Children   release:2019-04-26      tracks:1        artists: MATTN Klaas Roland Clark
Album[1]:       1tqc5vQIH7W6Hj74ZMbv9F  name:Lone Wolves        release:2019-07-19      tracks:1        artists: MATTN Paris Hilton
Album[2]:       6BUZhmN9MwcjhOFY7PPTLo  name:Cafe Del Mar 2016  release:2016-05-18      tracks:6        artists: MATTN Futuristic Polar Bears
Album[3]:       3cVeSfbkyYprXAmN30Kaup  name:Jungle Fever       release:2018-12-26      tracks:2        artists: MATTN
Album[4]:       2Wjd9pXMbXHGh9naXxWEjH  name:Late       release:2018-05-11      tracks:1        artists: MATTN HIDDN
Album[5]:       1PlDs0Ag56eGZ2lUveIjL7  name:Lone Wolves (Wekho Remix)  release:2019-08-30      tracks:1        artists: MATTN Paris Hilton Wekho
Album[6]:       54XzLf4USoeXfhCH2FvdWp  name:Throne     release:2018-10-03      tracks:1        artists: MATTN Futuristic Polar Bears Olly
Album[7]:       2LfIarjOBTsY7TWg0EdAlg  name:Don't Say A Word   release:2019-04-03      tracks:1        artists: MATTN
Album[8]:       2oGlQgQOIQM4g3dx30KPrh  name:Lone Wolves (Gaillard Remix)       release:2019-08-16      tracks:1        artists: MATTN Paris Hilton Gaillard
Album[9]:       2I6UfGH1bA7YFjtwMoX1ur  name:How We Roll        release:2016-06-27      tracks:1        artists: MATTN 2 Faced Funks
Album[10]:      2XSVOKgYiq4V37n4pQhYPX  name:Let The Song Play  release:2017-04-26      tracks:1        artists: MATTN Magic Wand
Album[11]:      0wXUxq3spEnZ1khPo0qLbB  name:Children   release:2019-05-29      tracks:2        artists: MATTN Klaas Roland Clark
Album[12]:      3A7fHCZfvh49Fpm96s2AkE  name:How We Roll        release:2016-07-20      tracks:1        artists: MATTN 2 Faced Funks
Album[13]:      40hhweFSHpwlFImxj8qrcO  name:Late       release:2018-06-06      tracks:2        artists: MATTN & HIDDN
> spoc list device
Device[0]:      eb2f5b8b312c7b67b23aba0586f20e3860286de3        name:mbp2016late        type:Computer   vol:100%        
Device[1]:      ffc16eaec9733d7ddebdcd5c7ae737dbc6b79fd9        name:HUAWEI Mate 9      type:Smartphone vol:100%        
> spoc tesujiro$ ./spoc play ffc16eaec9733d7ddebdcd5c7ae737dbc6b79fd9
> spoc tesujiro$ ./spoc pause ffc16eaec9733d7ddebdcd5c7ae737dbc6b79fd9
> spoc tesujiro$ ./spoc play next ffc16eaec9733d7ddebdcd5c7ae737dbc6b79fd9
> 
```
