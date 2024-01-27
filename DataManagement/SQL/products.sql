SELECT 

    DISTINCT MBRD, 
    COUNT(MBRD)  AS mbrd_count, 
    SUM(QTY) AS quantity 

FROM Products 
GROUP BY MBRD;