package bamboleo

type bamboleoTest struct {
	input    string
	expected entity
}

var stringTestCases = []bamboleoTest{
	{
		input: `{"#",acf6192e-81ca-46ef-93a6-5a6968b78663,
		{9,
			{3,
				{0,"Фрукт",{"Pattern",{"S"}},"Имя фрукта",0},
				{1,"Цвет",{"Pattern",{"S"}},"Цвет фрукта",0},
				{2,"Вес",{"Pattern",{"N",12,3,0}},"Вес фрукта",0}
			},
			{2,3,0,0,1,1,2,2,
				{1,2,
					{2,0,3,{"S","Яблоко"},{"S","Зелёный"},{"N",200},0},
					{2,1,3,{"S","Банан"},{"S","Желтый"},{"N",150},0}
				},
			2,1},
			{0,0}
		}
		}`,
		// expected: "PNG",
	},
	{
		input: `{"#",acf6192e-81ca-46ef-93a6-5a6968b78663,
		{9,
			{2,
				{0,"Артикул",{""},"",0},
				{1,"Цена",{""},"",0}
			},
		
			{2,2,0,0,1,1,
				{1,3,
					{2,0,2, {"S","Часы"}, {"N",750.99}, 0},
					{2,1,2, {"S","Кеды"}, {"N",150.99}, 0},
					{2,2,2, {"S","Шапка"}, {"N",200}, 0}
				},
			1,2},
		
			{0,0}
		}}`,
		// expected: "PNG",
	},
}
