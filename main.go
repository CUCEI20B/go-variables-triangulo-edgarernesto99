package main

import "fmt"

func main()  {
	var (
		base uint;
		altura uint;
		area uint;
	)
	fmt.Scan(&base);
	fmt.Scan(&altura);

	area = base*altura/2
	fmt.Println(area);
}