select
    t.team_id as team_id,
    t.team_name as team_name,
    case when sum(b.po) > 0 then sum(b.po)
         else 0 
    end as num_points
from(
    select
            a.host_team as ti, 
            sum(a.ho) as po
    from (
        select 
            host_team,
            case when host_goals > guest_goals then 3
                 when host_goals < guest_goals then 0
                 else 1
            end as ho
        from matches 
    ) a
    group by a.host_team
    union all
    select
            a.guest_team as ti,
            sum(a.gu) as po
    from (
        select 
            
            guest_team,
             case when host_goals < guest_goals then 3
                 when host_goals > guest_goals then 0
                 else 1
            end as gu
        from matches 
    ) a
    group by a.guest_team
) b right outer join teams t on b.ti  = t.team_id
group by t.team_id, b.ti, t.team_name
order by num_points desc, team_id asc