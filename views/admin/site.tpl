<div class="article">
    <form action="" method="post" class="login">
    {{ .Form | renderform }}
    Zone:
    <select name="country" id="">
        {{ range $item := .Country }}
            <option value="{{$item.ID}}">{{$item.Name}}</option>
        {{end}}
    </select>
    <label> Category:
        <input name="category" value="{{.Site.Category}}" type="text">
    </label>
    <button>add site</button>
</form>
</div>