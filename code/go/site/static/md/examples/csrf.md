## CSRF

## Explanation

Sometimes a backend framework need to set a header.  Normally you should be using cookies to be more secure, but it depends on your backend.

```html
<div data-header="{'x-csrf-token':'/Svi7DzhybrN+mDfI0zpReDj31ZZpp7GFp5KC6yMvGKer5OmslH1fpYDtAfsTwmfH+yLy7ghTAVHiRcjDz8XAQ=='}">
    <button data-on-click="$post('/examples/csrf/data')">Send update</button>
    <div>
        <label>Response</label>
        <div id="responses"></div>
    </div>
</div>
```

## Demo

<div id="update_me" data-on-load="$get('/examples/csrf/data')">Update Me</div>
