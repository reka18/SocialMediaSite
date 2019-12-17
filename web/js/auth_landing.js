
$( window ).on("load", async function()
{

    getPosts().then(function (postData)
    {
        let cardViews = document.getElementById('grid').innerHTML;

        if(!postData)
        {
            return;
        }

        if(postData.length)
        {
            for(let x = 0; x < postData.length ; x++)
            {
                const postObject = new UserPostData(postData[x]);
                cardViews += generate_user_card(postObject);

            }
            document.getElementById('grid').innerHTML = cardViews;
        }

        else
        {}


    });


});

$( window ).on("load", async function()
{

    getThisUser().then(function (userData)
    {

        if(!userData)
        {
            return;
        }

        if(userData.length)
        {
            const myInfoView = new UserData(userData);

            document.getElementById('myInfo').innerHTML = myInfoView;
        }

        else
        {}


    });


});


function updatePosts(postsData)
{
    console.log(postsData);
    const postObject = new UserPostData(postsData[0]);
    let newCard = generate_user_card(postObject);
    $('#grid').prepend(newCard);
}

function userLogout()
{
    window.location.href = "logout";
}

function showUserSearchResults(searchResults)
{
    document.getElementById('searchCardContainer').innerHTML="";
    let jsonSearchData = JSON.parse(searchResults);
    let searchCardViews = '';


    jsonSearchData.forEach( user =>
        {
            console.log(user);
            searchCardViews += search_results_card_generator(user);
        }
    );

    $('#searchCardContainer').prepend(searchCardViews);
    $('#exampleModalLong').modal('toggle');
}



