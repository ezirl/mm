<div class="article">
    <div class="title-page"><h1>Редактировать</h1></div>
<form action="/kosmo/site/edit" method="post" class="login">
    <input name="id" type="hidden"  value="{{.Site.ID}}">
    Name: <input name="name" type="text" value="{{.Site.Name}}">
    </br>Link: <input name="link" type="text" value="{{.Site.Link}}">
    </br>Rss: <input name="rss" type="text" value="{{.Site.Rss}}">
    </br> Ico: <input name="ico" type="text" value="{{.Site.Ico}}">
    <select name="country" id="">
                <option selected value="{{.Site.Country.ID}}">{{.Site.Country.Name}}<option>
                {{ range $item := .Country }}
                    <option value="{{$item.ID}}">{{$item.Name}}</option>
                {{end}}
            </select>
    Category: <input name="category" value="{{.Site.Category}}" type="text">
    <button>update site</button>
</form>
</div>