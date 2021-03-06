package migration

var SchemaVDE = `
		DROP TABLE IF EXISTS "%[1]d_vde_members";
		CREATE TABLE "%[1]d_vde_members" (
			"id" bigint NOT NULL DEFAULT '0',
			"member_name"	varchar(255) NOT NULL DEFAULT '',
			"image_id"	bigint,
			"member_info" jsonb
		);
		ALTER TABLE ONLY "%[1]d_vde_members" ADD CONSTRAINT "%[1]d_vde_members_pkey" PRIMARY KEY ("id");

		INSERT INTO "%[1]d_vde_members" ("id", "member_name") VALUES('%[2]d', 'founder');
		INSERT INTO "%[1]d_vde_members" ("id", "member_name") VALUES('4544233900443112470', 'guest');

		DROP TABLE IF EXISTS "%[1]d_vde_languages"; CREATE TABLE "%[1]d_vde_languages" (
		"id" bigint  NOT NULL DEFAULT '0',
		"name" character varying(100) NOT NULL DEFAULT '',
		"res" text NOT NULL DEFAULT ''
	  );
	  ALTER TABLE ONLY "%[1]d_vde_languages" ADD CONSTRAINT "%[1]d_vde_languages_pkey" PRIMARY KEY (id);
	  CREATE INDEX "%[1]d_vde_languages_index_name" ON "%[1]d_vde_languages" (name);
	  
	  DROP TABLE IF EXISTS "%[1]d_vde_menu"; CREATE TABLE "%[1]d_vde_menu" (
		  "id" bigint  NOT NULL DEFAULT '0',
		  "name" character varying(255) UNIQUE NOT NULL DEFAULT '',
		  "title" character varying(255) NOT NULL DEFAULT '',
		  "value" text NOT NULL DEFAULT '',
		  "conditions" text NOT NULL DEFAULT ''
	  );
	  ALTER TABLE ONLY "%[1]d_vde_menu" ADD CONSTRAINT "%[1]d_vde_menu_pkey" PRIMARY KEY (id);
	  CREATE INDEX "%[1]d_vde_menu_index_name" ON "%[1]d_vde_menu" (name);


	  INSERT INTO "%[1]d_vde_menu" ("id","name","title","value","conditions") VALUES('2','admin_menu','Admin menu','MenuItem(
    Icon: "icon-screen-desktop",
    Page: "interface",
    Vde: "true",
    Title: "Interface"
)
MenuItem(
    Icon: "icon-docs",
    Page: "tables",
    Vde: "true",
    Title: "Tables"
)
MenuItem(
    Icon: "icon-briefcase",
    Page: "contracts",
    Vde: "true",
    Title: "Smart Contracts"
)
MenuItem(
    Icon: "icon-settings",
    Page: "parameters",
    Vde: "true",
    Title: "Ecosystem parameters"
)
MenuItem(
    Icon: "icon-globe",
    Page: "languages",
    Vde: "true",
    Title: "Language resources"
)
MenuItem(
    Icon: "icon-cloud-upload",
    Page: "import",
    Vde: "true",
    Title: "Import"
)
MenuItem(
    Icon: "icon-cloud-download",
    Page: "export",
    Vde: "true",
    Title: "Export"
)','true');

	  DROP TABLE IF EXISTS "%[1]d_vde_pages"; CREATE TABLE "%[1]d_vde_pages" (
		  "id" bigint  NOT NULL DEFAULT '0',
		  "name" character varying(255) UNIQUE NOT NULL DEFAULT '',
		  "value" text NOT NULL DEFAULT '',
		  "menu" character varying(255) NOT NULL DEFAULT '',
		  "conditions" text NOT NULL DEFAULT '',
		  "validate_count" bigint NOT NULL DEFAULT '1',
		  "app_id" bigint NOT NULL DEFAULT '0',
		  "validate_mode" character(1) NOT NULL DEFAULT '0'
	  );
	  ALTER TABLE ONLY "%[1]d_vde_pages" ADD CONSTRAINT "%[1]d_vde_pages_pkey" PRIMARY KEY (id);
	  CREATE INDEX "%[1]d_vde_pages_index_name" ON "%[1]d_vde_pages" (name);

	  INSERT INTO "%[1]d_vde_pages" ("id","name","value","menu","conditions") VALUES('2','admin_index','','admin_menu','true');

	  DROP TABLE IF EXISTS "%[1]d_vde_blocks"; CREATE TABLE "%[1]d_vde_blocks" (
		  "id" bigint  NOT NULL DEFAULT '0',
		  "name" character varying(255) UNIQUE NOT NULL DEFAULT '',
		  "value" text NOT NULL DEFAULT '',
		  "conditions" text NOT NULL DEFAULT ''
	  );
	  ALTER TABLE ONLY "%[1]d_vde_blocks" ADD CONSTRAINT "%[1]d_vde_blocks_pkey" PRIMARY KEY (id);
	  CREATE INDEX "%[1]d_vde_blocks_index_name" ON "%[1]d_vde_blocks" (name);
	  
	  DROP TABLE IF EXISTS "%[1]d_vde_signatures"; CREATE TABLE "%[1]d_vde_signatures" (
		  "id" bigint  NOT NULL DEFAULT '0',
		  "name" character varying(100) NOT NULL DEFAULT '',
		  "value" jsonb,
		  "conditions" text NOT NULL DEFAULT ''
	  );
	  ALTER TABLE ONLY "%[1]d_vde_signatures" ADD CONSTRAINT "%[1]d_vde_signatures_pkey" PRIMARY KEY (name);
	  
	  CREATE TABLE "%[1]d_vde_contracts" (
	  "id" bigint NOT NULL  DEFAULT '0',
	  "name" text NOT NULL DEFAULT '',
	  "value" text  NOT NULL DEFAULT '',
	  "conditions" text  NOT NULL DEFAULT ''
	  );
	  ALTER TABLE ONLY "%[1]d_vde_contracts" ADD CONSTRAINT "%[1]d_vde_contracts_pkey" PRIMARY KEY (id);
	  
	  DROP TABLE IF EXISTS "%[1]d_vde_parameters";
	  CREATE TABLE "%[1]d_vde_parameters" (
	  "id" bigint NOT NULL  DEFAULT '0',
	  "name" varchar(255) UNIQUE NOT NULL DEFAULT '',
	  "value" text NOT NULL DEFAULT '',
	  "conditions" text  NOT NULL DEFAULT ''
	  );
	  ALTER TABLE ONLY "%[1]d_vde_parameters" ADD CONSTRAINT "%[1]d_vde_parameters_pkey" PRIMARY KEY ("id");
	  CREATE INDEX "%[1]d_vde_parameters_index_name" ON "%[1]d_vde_parameters" (name);
	  
	  INSERT INTO "%[1]d_vde_parameters" ("id","name", "value", "conditions") VALUES 
	  ('1','founder_account', '%[2]d', 'ContractConditions("MainCondition")'),
	  ('2','new_table', 'ContractConditions("MainCondition")', 'ContractConditions("MainCondition")'),
	  ('3','new_column', 'ContractConditions("MainCondition")', 'ContractConditions("MainCondition")'),
	  ('4','changing_tables', 'ContractConditions("MainCondition")', 'ContractConditions("MainCondition")'),
	  ('5','changing_language', 'ContractConditions("MainCondition")', 'ContractConditions("MainCondition")'),
	  ('6','changing_signature', 'ContractConditions("MainCondition")', 'ContractConditions("MainCondition")'),
	  ('7','changing_page', 'ContractConditions("MainCondition")', 'ContractConditions("MainCondition")'),
	  ('8','changing_menu', 'ContractConditions("MainCondition")', 'ContractConditions("MainCondition")'),
	  ('9','changing_contracts', 'ContractConditions("MainCondition")', 'ContractConditions("MainCondition")'),
	  ('10','stylesheet', 'body { 
		/* You can define your custom styles here or create custom CSS rules */
	  }', 'ContractConditions("MainCondition")'),
	  ('11','changing_blocks', 'ContractConditions("MainCondition")', 'ContractConditions("MainCondition")');

	  DROP TABLE IF EXISTS "%[1]d_vde_cron";
	  CREATE TABLE "%[1]d_vde_cron" (
		  "id"        bigint NOT NULL DEFAULT '0',
		  "owner"	  bigint NOT NULL DEFAULT '0',
		  "cron"      varchar(255) NOT NULL DEFAULT '',
		  "contract"  varchar(255) NOT NULL DEFAULT '',
		  "counter"   bigint NOT NULL DEFAULT '0',
		  "till"      timestamp NOT NULL DEFAULT timestamp '1970-01-01 00:00:00',
		  "conditions" text  NOT NULL DEFAULT ''
	  );
	  ALTER TABLE ONLY "%[1]d_vde_cron" ADD CONSTRAINT "%[1]d_vde_cron_pkey" PRIMARY KEY ("id");

		DROP TABLE IF EXISTS "%[1]d_vde_binaries";
		CREATE TABLE "%[1]d_vde_binaries" (
			"id" bigint NOT NULL DEFAULT '0',
			"app_id" bigint NOT NULL DEFAULT '1',
			"member_id" bigint NOT NULL DEFAULT '0',
			"name" varchar(255) NOT NULL DEFAULT '',
			"data" bytea NOT NULL DEFAULT '',
			"hash" varchar(32) NOT NULL DEFAULT '',
			"mime_type" varchar(255) NOT NULL DEFAULT ''
		);
		ALTER TABLE ONLY "%[1]d_vde_binaries" ADD CONSTRAINT "%[1]d_vde_binaries_pkey" PRIMARY KEY (id);
		CREATE UNIQUE INDEX "%[1]d_vde_binaries_index_app_id_member_id_name" ON "%[1]d_vde_binaries" (app_id, member_id, name);

	  CREATE TABLE "%[1]d_vde_tables" (
	  "id" bigint NOT NULL  DEFAULT '0',
	  "name" varchar(100) UNIQUE NOT NULL DEFAULT '',
	  "permissions" jsonb,
	  "columns" jsonb,
	  "conditions" text  NOT NULL DEFAULT '',
	  "app_id" bigint NOT NULL DEFAULT '1'
	  );
	  ALTER TABLE ONLY "%[1]d_vde_tables" ADD CONSTRAINT "%[1]d_vde_tables_pkey" PRIMARY KEY ("id");
	  CREATE INDEX "%[1]d_vde_tables_index_name" ON "%[1]d_vde_tables" (name);
	  
	  INSERT INTO "%[1]d_vde_tables" ("id", "name", "permissions","columns", "conditions") VALUES ('1', 'contracts', 
			  '{"insert": "ContractConditions(\"MainCondition\")", "update": "ContractConditions(\"MainCondition\")", 
				"new_column": "ContractConditions(\"MainCondition\")"}',
			  '{"name": "false",
				"value": "ContractConditions(\"MainCondition\")",
				"conditions": "ContractConditions(\"MainCondition\")"}', 'ContractAccess("EditTable")'),
			  ('2', 'languages', 
			  '{"insert": "ContractConditions(\"MainCondition\")", "update": "ContractConditions(\"MainCondition\")", 
				"new_column": "ContractConditions(\"MainCondition\")"}',
			  '{ "name": "ContractConditions(\"MainCondition\")",
				"res": "ContractConditions(\"MainCondition\")",
				"conditions": "ContractConditions(\"MainCondition\")"}', 'ContractAccess("EditTable")'),
			  ('3', 'menu', 
			  '{"insert": "ContractConditions(\"MainCondition\")", "update": "ContractConditions(\"MainCondition\")", 
				"new_column": "ContractConditions(\"MainCondition\")"}',
			  '{"name": "ContractConditions(\"MainCondition\")",
		  "value": "ContractConditions(\"MainCondition\")",
		  "conditions": "ContractConditions(\"MainCondition\")"
			  }', 'ContractAccess("EditTable")'),
			  ('4', 'pages', 
			  '{"insert": "ContractConditions(\"MainCondition\")", "update": "ContractConditions(\"MainCondition\")", 
				"new_column": "ContractConditions(\"MainCondition\")"}',
			  '{"name": "ContractConditions(\"MainCondition\")",
		  "value": "ContractConditions(\"MainCondition\")",
		  "menu": "ContractConditions(\"MainCondition\")",
		  "conditions": "ContractConditions(\"MainCondition\")",
		  "validate_count": "ContractConditions(\"MainCondition\")",
		  "validate_mode": "ContractConditions(\"MainCondition\")",
		  "app_id": "ContractConditions(\"MainCondition\")"
			  }', 'ContractAccess("EditTable")'),
			  ('5', 'blocks', 
			  '{"insert": "ContractConditions(\"MainCondition\")", "update": "ContractConditions(\"MainCondition\")", 
				"new_column": "ContractConditions(\"MainCondition\")"}',
			  '{"name": "ContractConditions(\"MainCondition\")",
		  "value": "ContractConditions(\"MainCondition\")",
		  "conditions": "ContractConditions(\"MainCondition\")"
			  }', 'ContractAccess("EditTable")'),
			  ('6', 'signatures', 
			  '{"insert": "ContractConditions(\"MainCondition\")", "update": "ContractConditions(\"MainCondition\")", 
				"new_column": "ContractConditions(\"MainCondition\")"}',
			  '{"name": "ContractConditions(\"MainCondition\")",
		  "value": "ContractConditions(\"MainCondition\")",
		  "conditions": "ContractConditions(\"MainCondition\")"
			  }', 'ContractAccess("EditTable")'),
			  ('7', 'cron',
				'{"insert": "ContractConditions(\"MainCondition\")", "update": "ContractConditions(\"MainCondition\")",
				  "new_column": "ContractConditions(\"MainCondition\")"}',
				'{"owner": "ContractConditions(\"MainCondition\")",
				"cron": "ContractConditions(\"MainCondition\")",
				"contract": "ContractConditions(\"MainCondition\")",
				"counter": "ContractConditions(\"MainCondition\")",
				"till": "ContractConditions(\"MainCondition\")",
				  "conditions": "ContractConditions(\"MainCondition\")"
				}', 'ContractConditions("MainCondition")'),
			  ('8', 'binaries',
				'{"insert": "ContractConditions(\"MainCondition\")", "update": "ContractConditions(\"MainCondition\")",
					"new_column": "ContractConditions(\"MainCondition\")"}',
				'{"app_id": "ContractConditions(\"MainCondition\")",
					"member_id": "ContractConditions(\"MainCondition\")",
					"name": "ContractConditions(\"MainCondition\")",
					"data": "ContractConditions(\"MainCondition\")",
					"hash": "ContractConditions(\"MainCondition\")",
					"mime_type": "ContractConditions(\"MainCondition\")"}',
					'ContractConditions("MainCondition")');
	  
	  INSERT INTO "%[1]d_vde_contracts" ("id", "name", "value", "conditions") VALUES 
	  ('1','MainCondition','contract MainCondition {
		conditions {
		  if EcosysParam("founder_account")!=$key_id
		  {
			warning "Sorry, you do not have access to this action."
		  }
		}
	  }', 'ContractConditions("MainCondition")'),
	  ('2','NewContract','contract NewContract {
		data {
			Value      string
			Conditions string
			Wallet         string "optional"
			TokenEcosystem int "optional"
			ApplicationId int "optional"
		}
		conditions {
			ValidateCondition($Conditions,$ecosystem_id)
			$walletContract = $key_id
			   if $Wallet {
				$walletContract = AddressToId($Wallet)
				if $walletContract == 0 {
				   error Sprintf("wrong wallet %%s", $Wallet)
				}
			}
			var list array
			list = ContractsList($Value)

			if Len(list) == 0 {
				error "must be the name"
			}

			var i int
			while i < Len(list) {
				if IsObject(list[i], $ecosystem_id) {
					warning Sprintf("Contract or function %%s exists", list[i] )
				}
				i = i + 1
			}

			$contract_name = list[0]
			if !$TokenEcosystem {
				$TokenEcosystem = 1
			} else {
				if !SysFuel($TokenEcosystem) {
					warning Sprintf("Ecosystem %%d is not system", $TokenEcosystem )
				}
			}
		}
		action {
			var root, id int
			root = CompileContract($Value, $ecosystem_id, $walletContract, $TokenEcosystem)
			id = DBInsert("contracts", "name,value,conditions, wallet_id, token_id,app_id",
				   $contract_name, $Value, $Conditions, $walletContract, $TokenEcosystem, $ApplicationId)
			FlushContract(root, id, false)
			$result = id
		}
		func rollback() {
			var list array
    		list = ContractsList($Value)
			var i int
			while i < Len(list) {
				RollbackContract(list[i])
				i = i + 1
			}
		}
		func price() int {
			return  SysParamInt("contract_price")
		}
	}', 'ContractConditions("MainCondition")'),
	  ('3','EditContract','contract EditContract {
		  data {
			  Id         int
			  Value      string "optional"
			  Conditions string "optional"
		  }

		  func onlyConditions() bool {
        	return $Conditions && !$Value
		  }
		  conditions {
			RowConditions("contracts", $Id, onlyConditions())
			if $Conditions {
	    		ValidateCondition($Conditions, $ecosystem_id)
			}

			var row array
			row = DBFind("contracts").Columns("id,value,conditions").WhereId($Id)
			if !Len(row) {
				error Sprintf("Contract %%d does not exist", $Id)
			}
			$cur = row[0]
			if $Value {
				var list, curlist array
				list = ContractsList($Value)
				curlist = ContractsList($cur["value"])
				if Len(list) != Len(curlist) {
					error "Contracts cannot be removed or inserted"
				}
				var i int
				while i < Len(list) {
					var j int
					var ok bool
					while j < Len(curlist) {
						if curlist[j] == list[i] {
							ok = true
							break
						}
						j = j + 1 
					}
					if !ok {
						error "Contracts or functions names cannot be changed"
					}
					i = i + 1
				}
			}
		  }
		  action {
			var root int
			var pars, vals array

			if $Value {
				root = CompileContract($Value, $ecosystem_id, 0, 0)
				pars[0] = "value"
				vals[0] = $Value
			}
			if $Conditions {
				pars[Len(pars)] = "conditions"
				vals[Len(vals)] = $Conditions
			}
			if Len(vals) > 0 {
				DBUpdate("contracts", $Id, Join(pars, ","), vals...)
			}
			if $Value {
			   FlushContract(root, $Id, false)
			}
		  }
	  }', 'ContractConditions("MainCondition")'),
	  ('4','NewParameter','contract NewParameter {
		  data {
			  Name string
			  Value string
			  Conditions string
		  }
		  conditions {
			  var ret array
			  ValidateCondition($Conditions, $ecosystem_id)
			  ret = DBFind("parameters").Columns("id").Where("name=?", $Name).Limit(1)
			  if Len(ret) > 0 {
				  warning Sprintf( "Parameter %%s already exists", $Name)
			  }
		  }
		  action {
			  $result = DBInsert("parameters", "name,value,conditions", $Name, $Value, $Conditions )
		  }
	  }', 'ContractConditions("MainCondition")'),
	  ('5','EditParameter','contract EditParameter {
		  data {
			  Id int
			  Value string
			  Conditions string
		  }
		  func onlyConditions() bool {
            	return $Conditions && !$Value
		  }
		  conditions {
			  RowConditions("parameters", $Id, onlyConditions())
			  ValidateCondition($Conditions, $ecosystem_id)
		  }
		  action {
			  DBUpdate("parameters", $Id, "value,conditions", $Value, $Conditions )
		  }
	  }', 'ContractConditions("MainCondition")'),
	  ('6', 'NewMenu','contract NewMenu {
		data {
			Name       string
			Value      string
			Title      string "optional"
			Conditions string
		}
		conditions {
			ValidateCondition($Conditions,$ecosystem_id)

			var row map
			row = DBRow("menu").Columns("id").Where("name = ?", $Name)

			if row {
				warning Sprintf( "Menu %%s already exists", $Name)
			}
		}
		action {
			DBInsert("menu", "name,value,title,conditions", $Name, $Value, $Title, $Conditions )
		}
		func price() int {
			return  SysParamInt("menu_price")
		}
	}', 'ContractConditions("MainCondition")'),
	  ('7','EditMenu','contract EditMenu {
		  data {
			  Id         int
			  Value      string "optional"
			  Title      string "optional"
			  Conditions string "optional"
		  }
		  
		func onlyConditions() bool {
        	return $Conditions && !$Value && !$Title
		}
	  	conditions {
		  RowConditions("menu", $Id, onlyConditions())
		  if $Conditions {
			  ValidateCondition($Conditions, $ecosystem_id)
		  }
	  	}
	  	action {
		  var pars, vals array
		  if $Value {
			  pars[0] = "value"
			  vals[0] = $Value
		  }
		  if $Title {
			  pars[Len(pars)] = "title"
			  vals[Len(vals)] = $Title
		  }
		  if $Conditions {
			  pars[Len(pars)] = "conditions"
			  vals[Len(vals)] = $Conditions
		  }
		  if Len(vals) > 0 {
			  DBUpdate("menu", $Id, Join(pars, ","), vals...)
		  }			
		}
	  }', 'ContractConditions("MainCondition")'),
	  ('8','AppendMenu','contract AppendMenu {
		data {
			Id     int
			Value  string
		}
		conditions {
			RowConditions("menu", $Id, false)
		}
		action {
			var row map
			row = DBRow("menu").Columns("value").WhereId($Id)
			DBUpdate("menu", $Id, "value", row["value"] + "\r\n" + $Value)
		}
	  }', 'ContractConditions("MainCondition")'),
	  ('9','NewPage','contract NewPage {
		data {
			Name       string
			Value      string
			Menu       string
			Conditions string
			ValidateCount int "optional"
			ApplicationId int "optional"
			ValidateMode int "optional"
		}
		func preparePageValidateCount(count int) int {
			var min, max int
			min = Int(EcosysParam("min_page_validate_count"))
			max = Int(EcosysParam("max_page_validate_count"))

			if count < min {
				count = min
			} else {
				if count > max {
					count = max
				}
			}

			return count
		}
		conditions {
			ValidateCondition($Conditions,$ecosystem_id)

			var row map
			row = DBRow("pages").Columns("id").Where("name = ?", $Name)

			if row {
				warning Sprintf( "Page %%s already exists", $Name)
			}

			$ValidateCount = preparePageValidateCount($ValidateCount)
		}
		action {
			DBInsert("pages", "name,value,menu,validate_count,conditions,app_id,validate_mode", 
				$Name, $Value, $Menu, $ValidateCount, $Conditions, $ApplicationId, $ValidateMode)
		}
		func price() int {
			return  SysParamInt("page_price")
		}
	}', 'ContractConditions("MainCondition")'),
	  ('10','EditPage','contract EditPage {
		  data {
			Id         int
			Value      string "optional"
			Menu      string "optional"
			Conditions string "optional"
			ValidateCount int "optional"
			ValidateMode  string "optional"
		  }
		  func onlyConditions() bool {
        	return $Conditions && !$Value && !$Menu
		  }
		  func preparePageValidateCount(count int) int {
			  var min, max int
			  min = Int(EcosysParam("min_page_validate_count"))
			  max = Int(EcosysParam("max_page_validate_count"))
	  
			  if count < min {
				  count = min
			  } else {
				  if count > max {
					  count = max
				  }
			  }
	  
			  return count
		  }					  
	  	conditions {
		  RowConditions("pages", $Id, onlyConditions())
		  if $Conditions {
			  ValidateCondition($Conditions, $ecosystem_id)
		  }
		  $ValidateCount = preparePageValidateCount($ValidateCount)
	  	}
	  	action {
			var pars, vals array
			if $Value {
				pars[0] = "value"
				vals[0] = $Value
			}
			if $Menu {
				pars[Len(pars)] = "menu"
				vals[Len(vals)] = $Menu
			}
			if $Conditions {
				pars[Len(pars)] = "conditions"
				vals[Len(vals)] = $Conditions
			}
			if $ValidateCount {
				pars[Len(pars)] = "validate_count"
				vals[Len(vals)] = $ValidateCount
			}
			if $ValidateMode {
				if $ValidateMode != "1" {
					$ValidateMode = "0"
				}
				pars[Len(pars)] = "validate_mode"
				vals[Len(vals)] = $ValidateMode
			}
			if Len(vals) > 0 {
				DBUpdate("pages", $Id, Join(pars, ","), vals...)
			}
	  	}		  
	  }', 'ContractConditions("MainCondition")'),
	  ('11','AppendPage','contract AppendPage {
		  data {
			  Id         int
			  Value      string
		  }
		  conditions {
			  RowConditions("pages", $Id, false)
		  }
		  action {
			  var row map
			  row = DBRow("pages").Columns("value").WhereId($Id)
			  DBUpdate("pages", $Id, "value", row["value"] + "\r\n" + $Value)
		  }
	  }', 'ContractConditions("MainCondition")'),
	  ('12','NewBlock','contract NewBlock {
		data {
			Name       string
			Value      string
			Conditions string
			ApplicationId int "optional"
		}
		conditions {
			ValidateCondition($Conditions,$ecosystem_id)

			var row map
			row = DBRow("blocks").Columns("id").Where("name = ?", $Name)

			if row {
				warning Sprintf( "Block %%s already exists", $Name)
			}
		}
		action {
			DBInsert("blocks", "name,value,conditions,app_id", $Name, $Value, $Conditions, $ApplicationId )
		}
	 }', 'ContractConditions("MainCondition")'),
	  ('13','EditBlock','contract EditBlock {
		data {
			Id         int
			Value      string "optional"
		  	Conditions string "optional"
		}
		
		func onlyConditions() bool {
			return $Conditions && !$Value
		}

	  	conditions {
			RowConditions("blocks", $Id, onlyConditions())
			if $Conditions {
				ValidateCondition($Conditions, $ecosystem_id)
			}
	  	}
	  	action {
		  var pars, vals array
		  if $Value {
			  pars[0] = "value"
			  vals[0] = $Value
		  }
		  if $Conditions {
			  pars[Len(pars)] = "conditions"
			  vals[Len(vals)] = $Conditions
		  }
		  if Len(vals) > 0 {
			  DBUpdate("blocks", $Id, Join(pars, ","), vals...)
		  }
		}
	  }', 'ContractConditions("MainCondition")'),
	  ('14','NewTable','contract NewTable {
		data {
			Name       string
			Columns      string
			Permissions string
			ApplicationId int "optional"
		}
		conditions {
			TableConditions($Name, $Columns, $Permissions)
		}
		action {
			CreateTable($Name, $Columns, $Permissions, $ApplicationId)
		}
		func rollback() {
			RollbackTable($Name)
		}
		func price() int {
			return  SysParamInt("table_price")
		}
	}', 'ContractConditions("MainCondition")'),
	  ('15','EditTable','contract EditTable {
		  data {
			  Name       string
			  Permissions string
		  }
		  conditions {
			  TableConditions($Name, "", $Permissions)
		  }
		  action {
			  PermTable($Name, $Permissions )
		  }
	  }', 'ContractConditions("MainCondition")'),
	  ('16','NewColumn','contract NewColumn {
		  data {
			  TableName   string
			  Name        string
			  Type        string
			  Permissions string
		  }
		  conditions {
			  ColumnCondition($TableName, $Name, $Type, $Permissions)
		  }
		  action {
			  CreateColumn($TableName, $Name, $Type, $Permissions)
		  }
	  }', 'ContractConditions("MainCondition")'),
	  ('17','EditColumn','contract EditColumn {
		  data {
			  TableName   string
			  Name        string
			  Permissions string
		  }
		  conditions {
			  ColumnCondition($TableName, $Name, "", $Permissions)
		  }
		  action {
			  PermColumn($TableName, $Name, $Permissions)
		  }
	  }', 'ContractConditions("MainCondition")'),
	  ('18','NewLang','contract NewLang {
		data {
			Name  string
			Trans string
			AppID int
		}
		conditions {
			EvalCondition("parameters", "changing_language", "value")
			var row array
			row = DBFind("languages").Columns("name").Where("name=? AND app_id=?", $Name, $AppID).Limit(1)
			if Len(row) > 0 {
				error Sprintf("The language resource %%s already exists", $Name)
			}
		}
		action {
			DBInsert("languages", "name,res,app_id", $Name, $Trans, $AppID)
			UpdateLang($AppID, $Name, $Trans)
		}
	}', 'ContractConditions("MainCondition")'),
	('19','EditLang','contract EditLang {
		data {
			Id    int
			Name  string
			Trans string
			AppID int
		}
		conditions {
			EvalCondition("parameters", "changing_language", "value")
		}
		action {
			DBUpdate("languages", $Id, "name,res,app_id", $Name, $Trans, $AppID)
			UpdateLang($AppID, $Name, $Trans)
		}
	}', 'ContractConditions("MainCondition")'),
	('20','Import','contract Import {
		data {
			Data string
		}
		conditions {
			$list = JSONDecode($Data)
		}
		func ImportList(row array, cnt string) {
			if !row {
				return
			}
			var i int
			while i < Len(row) {
				var idata map
				idata = row[i]
				if(cnt == "pages"){
					$ret_page = DBFind("pages").Columns("id").Where("name=$", idata["Name"])
					$page_id = One($ret_page, "id") 
					if ($page_id != nil){
						idata["Id"] = Int($page_id) 
						CallContract("EditPage", idata)
					} else {
						CallContract("NewPage", idata)
					}
				}
				if(cnt == "blocks"){
					$ret_block = DBFind("blocks").Columns("id").Where("name=$", idata["Name"])
					$block_id = One($ret_block, "id") 
					if ($block_id != nil){
						idata["Id"] = Int($block_id)
						CallContract("EditBlock", idata)
					} else {
						CallContract("NewBlock", idata)
					}
				}
				if(cnt == "menus"){
					$ret_menu = DBFind("menu").Columns("id,value").Where("name=$", idata["Name"])
					$menu_id = One($ret_menu, "id") 
					$menu_value = One($ret_menu, "value") 
					if ($menu_id != nil){
						idata["Id"] = Int($menu_id)
						idata["Value"] = Str($menu_value) + "\n" + Str(idata["Value"])
						CallContract("EditMenu", idata)
					} else {
						CallContract("NewMenu", idata)
					}
				}
				if(cnt == "parameters"){
					$ret_param = DBFind("parameters").Columns("id").Where("name=$", idata["Name"])
					$param_id = One($ret_param, "id")
					if ($param_id != nil){ 
						idata["Id"] = Int($param_id) 
						CallContract("EditParameter", idata)
					} else {
						CallContract("NewParameter", idata)
					}
				}
				if(cnt == "languages"){
					$ret_lang = DBFind("languages").Columns("id").Where("name=$", idata["Name"])
					$lang_id = One($ret_lang, "id")
					if ($lang_id != nil){
						CallContract("EditLang", idata)
					} else {
						CallContract("NewLang", idata)
					}
				}
				if(cnt == "contracts"){
					if IsObject(idata["Name"], $ecosystem_id){
					} else {
						CallContract("NewContract", idata)
					} 
				}
				if(cnt == "tables"){
					$ret_table = DBFind("tables").Columns("id").Where("name=$", idata["Name"])
					$table_id = One($ret_table, "id")
					if ($table_id != nil){	
					} else {
						CallContract("NewTable", idata)
					}
				}
				i = i + 1
			}
		}
		func ImportData(row array) {
			if !row {
				return
			}
			var i int
			while i < Len(row) {
				var idata map
				var list array
				var tblname, columns string
				idata = row[i]
				i = i + 1
				tblname = idata["Table"]
				columns = Join(idata["Columns"], ",")
				list = idata["Data"] 
				if !list {
					continue
				}
				var j int
				while j < Len(list) {
					var ilist array
					ilist = list[j]
					DBInsert(tblname, columns, ilist)
					j=j+1
				}
			}
		}
		action {
			ImportList($list["pages"], "pages")
			ImportList($list["blocks"], "blocks")
			ImportList($list["menus"], "menus")
			ImportList($list["parameters"], "parameters")
			ImportList($list["languages"], "languages")
			ImportList($list["contracts"], "contracts")
			ImportList($list["tables"], "tables")
			ImportData($list["data"])
		}
	}', 'ContractConditions("MainCondition")'),
	('21', 'NewCron','contract NewCron {
		data {
			Cron       string
			Contract   string
			Limit      int "optional"
			Till       string "optional date"
			Conditions string
		}
		conditions {
			ValidateCondition($Conditions,$ecosystem_id)
			ValidateCron($Cron)
		}
		action {
			if !$Till {
				$Till = "1970-01-01 00:00:00"
			}
			if !HasPrefix($Contract, "@") {
				$Contract = "@" + Str($ecosystem_id) + $Contract
			}
			$result = DBInsert("cron", "owner,cron,contract,counter,till,conditions",
				$key_id, $Cron, $Contract, $Limit, $Till, $Conditions)
			UpdateCron($result)
		}
	}', 'ContractConditions("MainCondition")'),
	('22','EditCron','contract EditCron {
		data {
			Id         int
			Contract   string
			Cron       string "optional"
			Limit      int "optional"
			Till       string "optional date"
			Conditions string
		}
		conditions {
			ConditionById("cron", true)
			ValidateCron($Cron)
		}
		action {
			if !$Till {
				$Till = "1970-01-01 00:00:00"
			}
			if !HasPrefix($Contract, "@") {
				$Contract = "@" + Str($ecosystem_id) + $Contract
			}
			DBUpdate("cron", $Id, "cron,contract,counter,till,conditions",
				$Cron, $Contract, $Limit, $Till, $Conditions)
			UpdateCron($Id)
		}
	}', 'ContractConditions("MainCondition")'),
	('23', 'UploadBinary', contract UploadBinary {
		data {
			Name  string
			Data  bytes "file"
			AppID int
			DataMimeType string "optional"
			MemberID int "optional"
		}
		conditions {
			$Id = Int(DBFind("binaries").Columns("id").Where("app_id = ? AND member_id = ? AND name = ?", $AppID, $MemberID, $Name).One("id"))
		}
		action {
			var hash string
			hash = MD5($Data)

			if $DataMimeType == "" {
				$DataMimeType = "application/octet-stream"
			}

			if $Id != 0 {
				DBUpdate("binaries", $Id, "data,hash,mime_type", $Data, hash, $DataMimeType)
			} else {
				$Id = DBInsert("binaries", "app_id,member_id,name,data,hash,mime_type", $AppID, $MemberID, $Name, $Data, hash, $DataMimeType)
			}

			$result = $Id
		}
	}', 'ContractConditions("MainCondition")');
	`
