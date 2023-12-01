package types

// Пример 1: https://www.plantuml.com/plantuml/umla/bPL1ZzCm48Nl_XLMVq1FY1D2ArX071QGeEA8lCqqjTQnaJEHa53-EmrcfqzJEosvrU_hU3plnRaRsMMUXc027TXoijdtHxQTpxHddw8v-exiGCCZvT6Odsx_v7gocs4QUMD_6tjw3ZxyM_d1DBBjqsOHokDYJEoMKlRRINmuKItj3Nr-QccVqbcUZTgskQFU3TTnDwtXRafsuqYPtwRO-ULt3NcSVEZUvPpo0usZwmcwlXJjTljQLxM5rQiBYQrN0UY_b96vsrVRhOn49YwhKbNcBI9Kd244cNhw1IbuNchwiSlXYsEc7Bzozh6tC1aiS3r7vqfGGYbByvKjPmD7sATuy3sSGTPMG_M1-6Fa5NNdL-Nt8Jb--MA5ADsk0FTJ_ljxhOaSQcff_hldupqTt1Iu0UtyTKgESWMK4rtXu60tD9Y3XcPpOUHjbv8qQ9-V5bqxIXEMiBGHRKURX6mr6D5lBYS6DZ14_gV779f8aMyj1ebeSSZK0t5fizniy2gtt0KMXg_i1GRVxgh32CDDg4n6JCsWwZJoykvNZ5weUkYqtwfq_KvGjXUodHbCW2vVloEgKIX_QDfUJmLkfICu6tFnhLW-g3XGK_7Rr0mpUkIuZXc4Jwy-eVpdZc8t3U4F
// пример 2: https://www.plantuml.com/plantuml/umla/hPR1YzGm5CRl-Il2dXIX9_6mY2owhdXO5NJnA7MQwOJD99Ac8CZyxtPIFqFxchuiqnvwo7lzvVlUUocvwKBjGtyqmgXzaC59hzj3a8tsQXUqiqA8N_NkiMwLl3hsNRYIVuGSdhqsIfuNMXSNudFGZHeYoZITnAAihNMXFhFaNMyZj8kGzBou7q3FdjErj-hpuzBw9xSOghjE-V3EsKQVjvc7V_RQDE-zT_vUTTrWAIN8anZVKtC_WhhS8RR_Dh2eKOIe8OI8tRdSiGtL_xslhubJMLLlu6YUFAkQa1ae6Y9VLnMYMOfSeS2FB2B18vD67LBbsv2W7i8M4z7iiyLC5RZFKUkLLd08I2KKQkQK3hfNhVgTplc_XQsEmOCtN-eGbBTVWzUsdKmrndgwzUR7C6d8dIdiIwew-JeMDgPEi-OUOtbJWNVExdMRAeoLpVwb7rJuQ0EypI8FEXkwCwuEhrvcekZNOl2szr5ORbDlvf5Hv7ST3hTgN_U6Y8fQcGn8Nak2ylKS6604URESWEcPztIRsSNWPZk11DuiG9WI6aM4Hs2SQ1WX3a87Zt8GBTJZ3OqYmYBeB5CEmlmWO0ReVuGqC_vAnlHfPkeKGynqdm779aaUdMGmG7IBGm9NB1NUYiJ26ij4CtagY6M2CN6bOb6sB99MYqSdR8bewg-yqjwiOOa8NZomPSBfs1TBnm2iqP7p93xvZZ51x8gJY7UBgkI64M-j8xTY9b88BY-E6cxdLURwMh2Pte22uRaKX1cX77gjvbdeBAVgHjccFvg_

type ClassModel struct {
	Name string // технически имена моделей и есть id. У меня id ассоциируется с цифрами. Хз как будет лучше
	Type classModelTypes
	//Visibility VisibilityKind // в plantuml отсутствуют параметры для  видимости для классов. Остается только стилями выделять
	style   string //#[color|back:color];header:color;line:color;line.[bold|dashed|dotted];text:color // стр 94
	extends string
	Field   *[]classFieldAndMethod
	Methods *[]classFieldAndMethod
}

type classModelTypes string

const (
	// Abstract      classModelTypes = "abstract" // закоментил, потому что конфликтует с classFieldAndMethodType.
	// 												 А так - это тоже самое, что и AbstractClass ниже
	AbstractClass classModelTypes = "abstract class"
	Annotation    classModelTypes = "annotation"
	Entity        classModelTypes = "entity"
	Enum          classModelTypes = "enum"
	Interface     classModelTypes = "interface"
	Protocol      classModelTypes = "protocol"
	Struct        classModelTypes = "struct"
	Exception     classModelTypes = "exception"
	Metaclass     classModelTypes = "metaclass"
	Stereotype    classModelTypes = "stereotype"
	// здесь отсутствуют CIRCLE, DIAMOND, тк они не поддерживают поля методов, а diamond даже название не поддреживает(хз вообще зачем они нужны, по идее это тип класса, но я с ним нихуя сделать не могу)
)

type classFieldAndMethod struct {
	Name       string
	Visibility VisibilityKind
	Type       classFieldAndMethodType
}

type classFieldAndMethodType string

const (
	Static   classFieldAndMethodType = "static"
	Abstract classFieldAndMethodType = "abstract"
	None     classFieldAndMethodType = ""
)
