CREATE VIEW statistics AS
WITH gral_credits AS (
    SELECT COUNT(*) total, SUM(invest) total_inv FROM credit_assigns
), success_credits AS (
    SELECT COUNT(*) total_sucess, SUM(invest) total_success_inv FROM credit_assigns
    WHERE status = 1
), fail_credits AS (
    SELECT COUNT(*) total_fails, SUM(invest) total_fail_inv FROM credit_assigns
    WHERE status = 0
)
SELECT gral_credits.total,
       gral_credits.total_inv,
       total_sucess,
       total_fails,
       success_credits.total_success_inv,
       ROUND(((success_credits.total_success_inv / gral_credits.total_inv::float)*100)::numeric, 2) avg_total_success_inv,
       ROUND(((total_fail_inv/total_inv::float)*100)::numeric, 2) avg_total_fail_inv
FROM gral_credits, success_credits, fail_credits;
