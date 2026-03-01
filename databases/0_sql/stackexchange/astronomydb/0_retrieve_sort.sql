-- 1
select
    p.title,
    p.score,
    case
        when score > 15 then 'High'
        when score <= 15
        and score >= 10 then 'Medium'
        else 'Low'
    end as popularity
from posts as p;

-- 2
select u.displayname, u.websiteurl
from users as u
order by
    case
        when u.websiteurl is null then 0
        else 1
    end desc,
    u.websiteurl;

-- 3
select avg(
        case
            when p.score is null then 0
            else p.score
        end
    )
from posts as p;

-- 4
select
    case
        when p.title is null then p.body
        else p.title
    end,
    case
        when p.posttypeid = 1 then 'Question'
        when p.posttypeid = 2 then 'Answer'
        else 'Other'
    end as post_type
from posts as p
    inner join posttypes as pt on p.posttypeid = pt.id;

-- 5
select v.postid, coalesce(p.title, 'No title') as title, sum(
        case
            when v.votetypeid = 2 then 2
            when v.votetypeid = 3 then -2
            else 0
        end
    ) as weight
from votes as v
    inner join posts as p on v.postid = p.id
group by
    v.postid,
    p.title
order by weight desc;

-- 6
select c.id, c.commenttext
from "comments" as c
order by random()
limit 5;

-- 7
select coalesce(
        coalesce(
            u.displayname, cast(u.accountid as varchar)
        ), 'Anonymous'
    ) as display_name
from users as u;

-- 8
select u.displayname, coalesce(
        u.lastaccessdate, u.creationdate
    ) as last_activity
from users as u
order by last_activity desc;

-- 9
select u.id, coalesce(u.location, 'Deep Space') from users as u;

-- 10
select p.title, p.viewcount, p.answercount
from posts as p
order by
    case
        when p.viewcount > 1000 then p.viewcount
        else p.answercount
    end desc;

-- 11