/* Embedding of Structs and Interfaces allows seamless Type composition and inheritance */

package basics

import "fmt"

type base struct {
    num int
}

func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

/* A container embeds a base. An embedding looks like a field without a name. */
type container struct {
    base // embedding
    str string
}

func RunEmbeddings() {
	co := container {
		/* When creating structs with literals, we have to initialize the embedding explicitly; 
		 * here the embedded type serves as the field name.*/
        base: base{ num: 5, },
        str: "container name",
	}
	
	// We can access the embedded type’s fields directly on co, e.g. co.num
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	
	// Alternatively, we can spell out the full path using the embedded type name
	fmt.Println("also num:", co.base.num)

	/* Since container embeds base, the methods of base also become methods of a container. 
	 * Here we invoke a method that was embedded from base directly on co. */
	fmt.Println("describe:", co.describe())


	/* Embedding structs with methods (such as embedding the base struct which implements a describe() method)
	 * may be used to bestow interface implementations onto other structs. */
	/* This `describer` interface not only declares the interface of the base struct
	 * (to which the describe() method was bound), but also the structs that embed this base struct. */
	type describer interface {
        describe() string
	}
	
	/* Here we see that a container now implements the describer interface because it embeds base.
	 * In other words, the container struct inherits the describe() method from the embedded base struct. */
	var d describer = co
    fmt.Println("describer of a container struct:", d.describe())
}