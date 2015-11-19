from pyquery import PyQuery as pq
import urllib
import codecs

import pandas as pd

LINKS = [
    'http://www.maccosmetics.com/products/13834/Products/Makeup/Eyes/Brow',
    'http://www.maccosmetics.com/products/13840/Products/Makeup/Eyes/Shadow',
    'http://www.maccosmetics.com/products/13835/Products/Makeup/Eyes/Eye-Palettes-Kits',
    'http://www.maccosmetics.com/products/14756/Products/Makeup/Eyes/Eye-Primer',
    'http://www.maccosmetics.com/products/13837/Products/Makeup/Eyes/Lash',
    'http://www.maccosmetics.com/products/13838/Products/Makeup/Eyes/Liner',
    'http://www.maccosmetics.com/products/13839/Products/Makeup/Eyes/Mascara',
    'http://www.maccosmetics.com/products/13825/Products/Skincare/Primers',
    'http://www.maccosmetics.com/products/14771/Products/Skincare/BB-CC',
    'http://www.maccosmetics.com/products/13824/Products/Skincare/Moisturizers',
    'http://www.maccosmetics.com/products/13825/Products/Skincare/Primers',
    'http://www.maccosmetics.com/products/13826/Products/Skincare/Removers',
    'http://www.maccosmetics.com/products/13827/Products/Skincare/Travel-Size',
    'http://www.maccosmetics.com/products/13804/Products/Brushes-Tools/Brushes/Eye-Brushes',
    'http://www.maccosmetics.com/products/13802/Products/Brushes-Tools/Brushes/All-Brushes',
    'http://www.maccosmetics.com/products/13803/Products/Brushes-Tools/Brushes/Brush-Kits',
    'http://www.maccosmetics.com/products/13805/Products/Brushes-Tools/Brushes/Face-Brushes',
    'http://www.maccosmetics.com/products/13806/Products/Brushes-Tools/Brushes/Lip-Brushes',
    'http://www.maccosmetics.com/products/13815/Products/Fragrance/Turquatic',
    
    
]


results = []
for link in LINKS:
    print "processing " + link
    d = pq(url=link, opener=lambda url, **kwargs: urllib.urlopen(url).read())
    
    for tomb in d("div.grid--mpp__item").items():
        name = tomb(".product__name").html()
        description = tomb(".product__description-short").html()
        price = tomb(".product__price").html()
        
        color_divs = tomb(".shade-picker__color-name")
        colors = "; ".join([pq(pq(c)).text() for c in color_divs]) or "**No Colors Found**"
        
        if name:
            results.append({
                'name': name,
                'description': description,
                'price': price,
                'colors': colors,
            })


df = pd.DataFrame(results)
print(df.head(10))
df.to_csv("products_crawl.csv")