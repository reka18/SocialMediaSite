
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
                cardViews += generate_post_card(postObject);

            }
            document.getElementById('grid').innerHTML = cardViews;
        }
        else
        {}
    });
});

function refreshPosts()
{
    getPosts().then(function (postData)
    {
        let cardViews = document.getElementById('grid').innerHTML = '';
        if(!postData)
        {
            return;
        }
        if(postData.length)
        {
            for(let x = 0; x < postData.length ; x++)
            {
                const postObject = new UserPostData(postData[x]);
                cardViews += generate_post_card(postObject);

            }
            document.getElementById('grid').innerHTML = cardViews;
        }
        else
        {}
    });

}

async function loadUserData(userName)
{

    getUser(userName).then(function (userData)
    {

        if(!userData)
        {
            return;
        }

        let data = new UserData(userData);
        document.getElementById('userNameId').innerHTML = data.getUsername();
        document.getElementById('ageId').innerHTML = data.getAge();
        document.getElementById('firstNameId').innerHTML = data.getFirstname();
        document.getElementById('lastNameId').innerHTML = data.getLastname();
        document.getElementById('emailId').innerHTML = data.getEmail();
        document.getElementById('isPublicId').innerHTML = data.getIsPublic();
        document.getElementById('friendlyDateId').innerHTML = data.getFriendlyJoinDate();
        document.getElementById('isActiveId').innerHTML = data.getIsActive();
        document.getElementById('genderId').innerHTML = data.getGender();
        document.getElementById('myInfoAvatar').src = 'avatar?user=' + data.getUsername();
    });


}

function otherUserSettings()
{

    document.getElementById('modalDims').style.width = "350px";
    document.getElementById('changeAvatarButton').style.visibility = "hidden";

}

function myUserSettings()
{

    document.getElementById('modalDims').style.width = "500px";
    document.getElementById('myInfoAvatar').className = 'w-100';
    document.getElementById('changeAvatarButton').style.visibility = "visible";

}


function updatePosts(postsData)
{
    console.log(postsData);
    const postObject = new UserPostData(postsData[0]);
    let newCard = generate_post_card(postObject);
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



