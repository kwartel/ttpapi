ttpapi
=======

A Go wrapper for the 32P [JSON API](https://github.com/TtpCD/Gazelle/wiki/JSON-API-Documentation)


Install
-------

```
go get "github.com/kwartel/ttpapi"
```

Example
-------
```Go
    ttp, err := ttpapi.NewTtpAPI("https://32PURL/")
    if err != nil {
        log.Fatal(err)
    }
    err = ttp.Login("username","password")
    if err != nil {
        log.Fatal(err)
    }
    mailboxParams := url.Values{}
    mailboxParams.Set("type", "sentbox")
    mailbox, err := ttp.GetMailbox(mailboxParams)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(mailbox)

    conversation, err := ttp.GetConversation(mailbox.Messages[0].ConvID)
    if err != nil {
        log.Fatal(err)
    }      
    log.Println(conversation.Messages[0].Body)

    torrentSearchParams := url.Values{}
    torrentSearch, err := ttp.SearchTorrents("Tool",torrentSearchParams)
    if err != nil {
        log.Fatal(err)
    }

    downloadURL, err := ttp.CreateDownloadURL(torrentSearch.Results[0].Torrents[0].TorrentID)
    if err != nil {
        log.Fatal(downloadURL)
    }
    log.Println(downloadURL)
```
