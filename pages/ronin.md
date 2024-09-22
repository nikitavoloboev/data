## Queries
```js
// https://discord.com/channels/1099977902629589095/1258718939312160789/1258790367348916350
// Everything will be typed here as expected
const [categories, productsIncludingCategory, productsCount, product, ...categoriesNew] = await batch(() => [
    get.categories(),
    get.products.including(["category"]),
    count.products(),
    set.product({
        with: {  id: "asdasd" },
        to: { active: false },
        including: ["category"],
    }),

    ...localCategories.map((category) =>
        create.category.with({
            id: category.id,
            name: category.name,
        })
    ),
])
```
