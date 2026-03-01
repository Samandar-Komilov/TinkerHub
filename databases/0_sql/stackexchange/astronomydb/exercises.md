# Astronomy Stack Exchange Dataset Exercises

### Retrieval and Sorting

##### Conditional Logic & Selection

1. **Category Mapping**: Select the `Title` and `Score` from `Posts`. Use a `CASE` expression to create a new column `Popularity` where scores > 10 are 'High', scores between 1 and 10 are 'Medium', and scores <= 0 are 'Low'.
2. **Conditional Sorting with NULLs**: Select `DisplayName` and `WebsiteUrl` from `Users`. Order the results such that users with a `WebsiteUrl` appear first, sorted alphabetically, followed by users without a URL.
3. **NULL-Safe Aggregation**: Calculate the average `Score` of all `Posts`. Use `CASE` within the `AVG` function to treat any `NULL` scores as 0 to ensure they are factored into the calculation.
4. **Post Type Labeling**: Join `Posts` and `PostTypes`. Use `CASE` to return the `Title` (or `Body` if `Title` is NULL) and a custom label: 'Question' for `PostTypeId` 1, 'Answer' for 2, and 'Other' for any other ID.
5. **Vote Impact**: Select `PostId` from `Votes`. Use a `CASE` expression to assign a numeric value `Weight`: 2 for 'UpMod' (VoteTypeId 2), -2 for 'DownMod' (VoteTypeId 3), and 0 for all others. Sum this `Weight` per `PostId`.

##### Returning Random Records

6. **Random Sampling**: Return 5 random `Body` snippets from the `Comments` table. In PostgreSQL, utilize `ORDER BY RANDOM()` combined with a `LIMIT` clause.

##### Transforming NULLs (COALESCE)

7. **Display Name Fallback**: Select `DisplayName` from `Users`. If `DisplayName` is NULL, return their `AccountId`. If both are NULL, return the string 'Anonymous'.
8. **Last Activity Heatmap**: Select `Title` and `LastEditDate`. Use `COALESCE` to return `LastEditDate`, but if that is NULL, return `CreationDate` to ensure every row has a timestamp representing "Last Known Activity".
9. **Location Defaults**: Select `Id` and `Location` from `Users`. Use `COALESCE` to replace all NULL locations with the string 'Deep Space'.

##### Sorting on Data-Dependent Keys

10. **Dynamic Ordering (Image 1 Methodology)**: Select `Title`, `ViewCount`, and `AnswerCount`. If `ViewCount` is greater than 1000, sort the result set by `ViewCount` descending; otherwise, sort by `AnswerCount` descending.
* *Note: Use the `CASE` expression directly in the `ORDER BY` clause.*


11. **Post Format Preference**: Select `Title` and `Tags`. Use a `CASE` in the `ORDER BY` to prioritize posts that contain the tag '<black-holes>' first, then sort by `CreationDate` for everything else.
12. **Reputation-Based Sorting**: Select `DisplayName` and `Reputation`. If a user's `Reputation` is over 5000, sort them by `DisplayName` (alphabetical). For users under 5000, sort by `Reputation` (highest first).
13. **Length-Based Priority**: Select `Text` from `Comments`. Sort the results based on the `PostId`. However, for `PostId` values divisible by 2, sort by the length of the `Text` (short to long). For odd `PostId` values, sort alphabetically.
14. **Conditional Null Placement**: Select `DisplayName` and `Location`. Sort the results so that users in 'United Kingdom' appear at the top, followed by users in 'USA', then everyone else sorted alphabetically, with NULL locations forced to the very bottom regardless of default behavior.

---

### Window Functions

Window functions perform calculations across a set of table rows that are related to the current row. Unlike aggregate functions, they do not cause rows to become grouped into a single output row; the rows retain their separate identities.

##### Ranking & Identification

1. **Global Rank**: Select `DisplayName` and `Reputation` from `Users`. Use `RANK()` to assign a position to each user based on their `Reputation` (highest first).
2. **Top Questions per Year**: From `Posts`, select the `Title`, `Score`, and the year of `CreationDate`. Use `DENSE_RANK()` to rank posts by `Score` descending, partitioned by the year.
3. **Row Identification**: Select `Id` and `Body` from `Comments`. Use `ROW_NUMBER()` to assign a unique sequential integer to each comment made by a specific `UserId`, ordered by `CreationDate`.

##### Running Totals & Moving Averages

4. **Cumulative Score**: Select `Id`, `CreationDate`, and `Score` from `Posts`. Calculate a running total of the `Score` for all posts, ordered chronologically by `CreationDate`.
5. **Monthly Growth**: Select the month of `CreationDate` and count the number of new `Users`. Use a window function to show the cumulative sum of new users month-over-month.
6. **Moving Average**: Select `Id` and `ViewCount` from `Posts`. Calculate a 5-row moving average of `ViewCount` (the current row plus the four preceding rows) ordered by `Id`.

##### Positional (Lead/Lag)

7. **Time Between Posts**: For each user in the `Posts` table, select the `Title` and `CreationDate`. Use `LAG()` to show the `CreationDate` of the *previous* post by that same user.
8. **Score Gap**: Select `Title` and `Score`. Use `LEAD()` to compare the current post's `Score` with the `Score` of the post immediately following it in terms of `ViewCount`.

##### Statistical Distributions

9. **Reputation Percentile**: Select `DisplayName` and `Reputation`. Use `PERCENT_RANK()` to determine the percentile ranking of each user’s reputation within the community.
10. **Quartile Buckets**: Divide the `Posts` table into 4 equal groups (quartiles) based on `ViewCount` using the `NTILE(4)` function. Return the `Title`, `ViewCount`, and its assigned quartile.