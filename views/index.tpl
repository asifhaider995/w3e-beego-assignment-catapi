{{template "layout.tpl" . }}

{{define "styles" }}
    <!--suppress ALL -->
    <style type="text/css">
        .main {
            display: flex;
            justify-content: center;
            align-items: center;
        }
        .centered {
            display: flex;
            flex-direction: column;
            align-items: center;
            justify-content: space-evenly;
            min-height: 500px;
        }
        .btns {
            display: flex;
            justify-content: space-evenly;
            width: 100%;
        }
        .btn {
            padding: 5px 10px;
            border-radius: 5px;
            border: 1px solid black;
            background: white;
            color: black;
        }
        .btn:hover {
            background: black;
            color: white;
            cursor: pointer;
        }
    </style>

{{end}}

{{define "content"}}
    <div class="main">
        <main class="centered">
            <h1>Do you like this cat?</h1>
            <div class="imgDiv">
                <img class="mainImg" src="/" alt="cat"/>
            </div>
            <div class="btns">
                <button class="btn" id="like">Vote up</button>
                <button class="btn" id="dislike">Vote down</button>
                <button class="btn" id="favorite">Favorite</button>
                <button class="btn" id="skip">Skip</button>
            </div>
        </main>
    </div>
{{end}}

{{define "website"}}
    {{.Website}}
{{end}}

{{define "email"}}
    {{.Email}}
{{end}}

{{define "scripts"}}
    <script>
        let val = JSON.parse({{.jsonStr}});
        console.log(val.url)
        console.log(val.id)
        const nextImg = () => {

        }
        u('.mainImg').attr('src', val.url).attr('width', 350).attr('height', 250);
        // lis.append('<img src={val.url} alt={"cat"}>')
        u('#like').on('click', () => {
            console.log("Liked", val.id)
            const url = "https://api.thecatapi.com/v1/votes"
            const api_key = "dcd005e2-35a8-45ec-843d-fcdfac3a4869"
            let xml = new XMLHttpRequest();
            // let fd = new FormData();
            // fd.append("image_id", val.id)
            // fd.append("value", 1)
            const pl = {
                "image_id": val.id,
                "value": 1
            }

            xml.onload = () => {
                console.log(xml.responseText)
                let xml2 = new XMLHttpRequest();

                xml2.onload = function () {
                    // console.log(JSON.parse(xml.response)[0].url)
                    u('.mainImg').attr('src', JSON.parse(xml2.response)[0].url).attr('width', 350).attr('height', 250);
                }
                xml2.open("GET", {{.api}})
                xml2.send()
            }
            xml.open("POST", url, true);
            xml.setRequestHeader('Access-Control-Allow-Origin', "*");
            xml.setRequestHeader('Content-Type', "application/json");
            xml.setRequestHeader('x-api-key', api_key);
            xml.send(JSON.stringify(pl))
        })
        u('#dislike').on('click', () => {
            console.log("Disliked", val.id)
            const url = "https://api.thecatapi.com/v1/votes"
            const api_key = "dcd005e2-35a8-45ec-843d-fcdfac3a4869"
            let xml = new XMLHttpRequest();
            const pl = {
                "image_id": val.id,
                "value": 0
            }

            xml.onload = () => {
                console.log(xml.responseText)
                let xml2 = new XMLHttpRequest();

                xml2.onload = function () {
                    // console.log(JSON.parse(xml.response)[0].url)
                    u('.mainImg').attr('src', JSON.parse(xml2.response)[0].url).attr('width', 350).attr('height', 250);
                }
                xml2.open("GET", {{.api}})
                xml2.send()
            }
            xml.open("POST", url, true);
            xml.setRequestHeader('Access-Control-Allow-Origin', "*");
            xml.setRequestHeader('Content-Type', "application/json");
            xml.setRequestHeader('x-api-key', api_key);
            xml.send(JSON.stringify(pl))
        })
        u('#favorite').on('click', () => {
            console.log("favorite", val)
            const url = "https://api.thecatapi.com/v1/favourites"
            const api_key = "dcd005e2-35a8-45ec-843d-fcdfac3a4869"
            let xml = new XMLHttpRequest();
            const pl = {
                "image_id": val.id,
            }

            xml.onload = () => {
                console.log(xml.responseText)
                let xml2 = new XMLHttpRequest();

                xml2.onload = function () {
                    // console.log(JSON.parse(xml.response)[0].url)
                    u('.mainImg').attr('src', JSON.parse(xml2.response)[0].url).attr('width', 350).attr('height', 250);
                }
                xml2.open("GET", {{.api}})
                xml2.send()
            }
            xml.open("POST", url, true);
            xml.setRequestHeader('Access-Control-Allow-Origin', "*");
            xml.setRequestHeader('Content-Type', "application/json");
            xml.setRequestHeader('x-api-key', api_key);
            xml.send(JSON.stringify(pl))

        })
        u('#skip').on('click', () => {
            console.log("Skipped")

            let xml = new XMLHttpRequest();

            xml.onload = function () {
                // console.log(JSON.parse(xml.response)[0].url)
                u('.mainImg').attr('src', JSON.parse(xml.response)[0].url).attr('width', 350).attr('height', 250);
            }
            xml.open("GET", {{.api}})
            xml.send()


        })
    </script>
{{end}}

