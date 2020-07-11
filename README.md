# govis

It is a tool written in golang to make screenshotting urls more reliable and consistent. It scans url for screenshot and Javascript files and presents the result in html form.

![help](https://i.ibb.co/b26QhpL/carbon-1.png)

# Feautres

* It also **captures all the js files** loaded by the domain.
* **Reliable** that other tools I have tested.
* **STDIN** support to help in workflows
* Saves info in js file which can be used as **json** containing 
 `{
    Url,
    Title,
    StatusCode,
    JsUrlsList
  }` with some tweaks
* Simple to use 
```bash 
    ▶ cat urls | govis
```

# Installation

```bash
   ▶ go get -u github.com/0xrishabh/govis

```
