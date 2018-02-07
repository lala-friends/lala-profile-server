package util

const SELECT_PERSON = "SELECT NAME, EMAIL, INTORUDUCE, IMAGE_URL, REP_COLOR	 ,BLOG, GITHUB, FACEBOOK FROM PERSON WHERE ID = ?"
const SELECT_PERSON_ALL = "SELECT ID, NAME, EMAIL, INTORUDUCE, IMAGE_URL, REP_COLOR	 ,BLOG, GITHUB, FACEBOOK FROM PERSON"
const SELECT_PROJECTS = "SELECT PROJECT_NAME, PERIOD, PERSONAL_ROLE, MAIN_OPERATOR, PROJECT_SUMMARY, RESPONSIBILITIES, USED_TECHNOLOGY, PRIMARY_ROLE, PROJECT_RESULT, LINKED_SITE FROM PROJECT WHERE PERSON_ID = ?"
const SELECT_PRODUCT_BY_PERSON = "SELECT pdt.Id, pdt.NAME, pdt.INTRODUCE, pdt.TECH, pdt.IMAGE_URL FROM PRODUCT pdt, MAP_PERSON_PRODUCT mpp, PERSON psn WHERE pdt.ID = mpp.PRODUCT_ID AND mpp.PERSON_ID = psn.ID AND psn.ID = ?"

const SELECT_PRODUCT_ALL = "SELECT pdt.ID, pdt.NAME, pdt.INTRODUCE, pdt.TECH, pdt.IMAGE_URL FROM PRODUCT pdt"
const SELECT_PERSON_BY_PRODUCT =
	"SELECT psn.ID, " +
		  "psn.NAME " +
		 ",psn.EMAIL "+
		 ",psn.INTORUDUCE "+
		 ",psn.IMAGE_URL "+
   		 ",psn.REP_COLOR "+
		 ",psn.BLOG "+
		 ",psn.GITHUB "+
		 ",psn.FACEBOOK "+
		 "FROM PERSON psn "+
		 ",(SELECT mpp.PERSON_ID AS mppPersonId "+
			"FROM MAP_PERSON_PRODUCT mpp "+
			"WHERE mpp.PRODUCT_ID = ? ) innerMpp "+
 		 "WHERE psn.ID = innerMpp.mppPersonId"