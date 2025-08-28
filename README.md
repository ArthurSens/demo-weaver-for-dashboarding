# demo-weaver-for-dashboarding

Generate all metrics and relevant attributes:

```bash
weaver registry generate  --registry=../semantic-conventions/model --templates=./templates go ./output
```

Generate metrics and relevant attributes in specific namespaces:

```bash
weaver registry generate  --registry=../semantic-conventions/model --templates=./templates go ./output --param included_namespaces="db,http"
```