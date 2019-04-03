# --------------------------------------------------------------------------------------------------
# prgn.R (https://github.com/stefanomozart/computational-statistics/R/prgn.R)
# copywrite 2019 - Stefano Mozart (stefanomozart@ieee.org)
#.This work is distributed under the GNU General Public Licence (https://www.gnu.org/licenses/gpl-3.0.en.html)
# --------------------------------------------------------------------------------------------------

LCG <- function(m, a, c, Xo) {
	return (((a * Xo) + c) %% m);
};

rLCG <- function(n, m, a, c, Xo) {
	X <- c()
	X[1] <- Xo
	for (i in 2:(n+1)) {
		X[i] <- LCG(m, a, c, X[i-1])
	}
	if (n==1) {
		return(X[i])
	}
	return(X);
};

