# govis

It is a tool written in golang to make screenshotting urls more reliable and consistent. It scans url for screenshot and Javascript files and presents the result in html form.

# Feautres

* It also **captures all the js files** loaded by the domain.
* **Reliable** that other tools I have tested.
* **STDIN** support to help in worlkflows
* Saves info in js file which can be used as **json** containing 
 `{
    Url,
    Title,
    StatusCode,
    JsUrlsList
  }` with some twi
* Simple to use 
```bash 
    cat urls | govis
```

# Installation

```bash
   go get github.com/0xrishabh/govis

```
