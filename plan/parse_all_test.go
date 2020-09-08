package plan

import (
	"fmt"
	. "github.com/pingcap/check"
)

func (s *parseTestSuite) TestParseAll(c *C) {
	ep := `
+------------------------------------------------+--------+------+------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| id                                             | count  | task | operator info                                                                                                                                                                                                                                                                                                                                                                                              |
+------------------------------------------------+--------+------+------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| Projection_31                                  | 62.83  | root | sdyx.r.resource_id, sdyx.r.status, sdyx.r.action_url, sdyx.r.resource_name, sdyx.r.parent_id, sdyx.r.url, sdyx.r.resources_icon_url, sdyx.r.sort_code, sdyx.r.sys_id, case count(*)           when 0             then 0           else 1 end, ifnull(ifnull(rr1.can_add,"1"), "1"), ifnull(ifnull(rr2.can_delete,"1"), "1"), ifnull(ifnull(rr3.can_other,"1"), "1"), ifnull(ifnull(rr4.can_look,"1"), "1") |
| └─Apply_33                                     | 62.83  | root | left outer join, inner:MaxOneRow_115                                                                                                                                                                                                                                                                                                                                                                       |
|   ├─Apply_35                                   | 62.83  | root | left outer join, inner:MaxOneRow_102                                                                                                                                                                                                                                                                                                                                                                       |
|   │ ├─Apply_37                                 | 62.83  | root | left outer join, inner:MaxOneRow_89                                                                                                                                                                                                                                                                                                                                                                        |
|   │ │ ├─Apply_39                               | 62.83  | root | left outer join, inner:MaxOneRow_76                                                                                                                                                                                                                                                                                                                                                                        |
|   │ │ │ ├─Projection_40                        | 62.83  | root | sdyx.r.resource_id, sdyx.r.status, sdyx.r.action_url, sdyx.r.resource_name, sdyx.r.parent_id, sdyx.r.url, sdyx.r.resources_icon_url, sdyx.r.sort_code, sdyx.r.sys_id, case(eq(6_col_0, 0), 0, 1)                                                                                                                                                                                                           |
|   │ │ │ │ └─Projection_41                      | 62.83  | root | sdyx.r.resource_id, sdyx.r.status, sdyx.r.action_url, sdyx.r.resource_name, sdyx.r.parent_id, sdyx.r.url, sdyx.r.resources_icon_url, sdyx.r.sort_code, sdyx.r.sys_id, ifnull(6_col_0, 0)                                                                                                                                                                                                                   |
|   │ │ │ │   └─HashLeftJoin_42                  | 62.83  | root | left outer join, inner:HashAgg_58, equal:[eq(sdyx.r.resource_id, sdyx.rr.resource_id)]                                                                                                                                                                                                                                                                                                                     |
|   │ │ │ │     ├─TableReader_45                 | 62.83  | root | data:Selection_44                                                                                                                                                                                                                                                                                                                                                                                          |
|   │ │ │ │     │ └─Selection_44                 | 62.83  | cop  | eq(sdyx.r.sys_id, "C"), ne(sdyx.r.status, "1")                                                                                                                                                                                                                                                                                                                                                             |
|   │ │ │ │     │   └─TableScan_43               | 235.00 | cop  | table:R, range:[-inf,+inf], keep order:false                                                                                                                                                                                                                                                                                                                                                               |
|   │ │ │ │     └─HashAgg_58                     | 6.00   | root | group by:col_2, funcs:count(col_0), firstrow(col_1)                                                                                                                                                                                                                                                                                                                                                        |
|   │ │ │ │       └─IndexLookUp_59               | 6.00   | root |                                                                                                                                                                                                                                                                                                                                                                                                            |
|   │ │ │ │         ├─IndexScan_56               | 19.00  | cop  | table:RR, index:USER_ROLE_ID, range:["RGxRpwZZJjfqW8ZRM7U","RGxRpwZZJjfqW8ZRM7U"], keep order:false                                                                                                                                                                                                                                                                                                        |
|   │ │ │ │         └─HashAgg_47                 | 6.00   | cop  | group by:sdyx.rr.resource_id, funcs:count(1), firstrow(sdyx.rr.resource_id)                                                                                                                                                                                                                                                                                                                                |
|   │ │ │ │           └─TableScan_57             | 19.00  | cop  | table:SYS_ROLES_RESOURCE, keep order:false                                                                                                                                                                                                                                                                                                                                                                 |
|   │ │ │ └─MaxOneRow_76                         | 1.00   | root |                                                                                                                                                                                                                                                                                                                                                                                                            |
|   │ │ │   └─Projection_77                      | 0.10   | root | ifnull(sdyx.rr1.can_add, "1")                                                                                                                                                                                                                                                                                                                                                                              |
|   │ │ │     └─IndexLookUp_84                   | 0.10   | root |                                                                                                                                                                                                                                                                                                                                                                                                            |
|   │ │ │       ├─IndexScan_81                   | 3.16   | cop  | table:RR1, index:RESOURCE_ID, range: decided by [eq(sdyx.rr1.resource_id, sdyx.r.resource_id)], keep order:false                                                                                                                                                                                                                                                                                           |
|   │ │ │       └─Selection_83                   | 0.10   | cop  | eq(sdyx.rr1.user_role_id, "RGxRpwZZJjfqW8ZRM7U")                                                                                                                                                                                                                                                                                                                                                           |
|   │ │ │         └─TableScan_82                 | 3.16   | cop  | table:SYS_ROLES_RESOURCE, keep order:false                                                                                                                                                                                                                                                                                                                                                                 |
|   │ │ └─MaxOneRow_89                           | 1.00   | root |                                                                                                                                                                                                                                                                                                                                                                                                            |
|   │ │   └─Projection_90                        | 0.10   | root | ifnull(sdyx.rr2.can_delete, "1")                                                                                                                                                                                                                                                                                                                                                                           |
|   │ │     └─IndexLookUp_97                     | 0.10   | root |                                                                                                                                                                                                                                                                                                                                                                                                            |
|   │ │       ├─IndexScan_94                     | 3.16   | cop  | table:RR2, index:RESOURCE_ID, range: decided by [eq(sdyx.rr2.resource_id, sdyx.r.resource_id)], keep order:false                                                                                                                                                                                                                                                                                           |
|   │ │       └─Selection_96                     | 0.10   | cop  | eq(sdyx.rr2.user_role_id, "RGxRpwZZJjfqW8ZRM7U")                                                                                                                                                                                                                                                                                                                                                           |
|   │ │         └─TableScan_95                   | 3.16   | cop  | table:SYS_ROLES_RESOURCE, keep order:false                                                                                                                                                                                                                                                                                                                                                                 |
|   │ └─MaxOneRow_102                            | 1.00   | root |                                                                                                                                                                                                                                                                                                                                                                                                            |
|   │   └─Projection_103                         | 0.10   | root | ifnull(sdyx.rr3.can_other, "1")                                                                                                                                                                                                                                                                                                                                                                            |
|   │     └─IndexLookUp_110                      | 0.10   | root |                                                                                                                                                                                                                                                                                                                                                                                                            |
|   │       ├─IndexScan_107                      | 3.16   | cop  | table:RR3, index:RESOURCE_ID, range: decided by [eq(sdyx.rr3.resource_id, sdyx.r.resource_id)], keep order:false                                                                                                                                                                                                                                                                                           |
|   │       └─Selection_109                      | 0.10   | cop  | eq(sdyx.rr3.user_role_id, "RGxRpwZZJjfqW8ZRM7U")                                                                                                                                                                                                                                                                                                                                                           |
|   │         └─TableScan_108                    | 3.16   | cop  | table:SYS_ROLES_RESOURCE, keep order:false                                                                                                                                                                                                                                                                                                                                                                 |
|   └─MaxOneRow_115                              | 1.00   | root |                                                                                                                                                                                                                                                                                                                                                                                                            |
|     └─Projection_116                           | 0.10   | root | ifnull(sdyx.rr4.can_look, "1")                                                                                                                                                                                                                                                                                                                                                                             |
|       └─IndexLookUp_123                        | 0.10   | root |                                                                                                                                                                                                                                                                                                                                                                                                            |
|         ├─IndexScan_120                        | 3.16   | cop  | table:RR4, index:RESOURCE_ID, range: decided by [eq(sdyx.rr4.resource_id, sdyx.r.resource_id)], keep order:false                                                                                                                                                                                                                                                                                           |
|         └─Selection_122                        | 0.10   | cop  | eq(sdyx.rr4.user_role_id, "RGxRpwZZJjfqW8ZRM7U")                                                                                                                                                                                                                                                                                                                                                           |
|           └─TableScan_121                      | 3.16   | cop  | table:SYS_ROLES_RESOURCE, keep order:false                                                                                                                                                                                                                                                                                                                                                                 |
+------------------------------------------------+--------+------+------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
`

	vers := []string{V2, V3, V4}
	for _, ver := range vers {
		fmt.Println("test version ", ver)
		_, err := ParseText("", ep, ver)
		c.Assert(err, IsNil)
	}
}