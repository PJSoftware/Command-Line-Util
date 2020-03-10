#!/perl -w

use strict;

print "Enter Season number: ";
my $S = <stdin>;
chomp $S;

print "Enter # of discs: ";
my $D = <stdin>;
chomp $D;

foreach my $season (1 .. $S) {

    mkdir("s$season",0);
    chdir("s$season");

    foreach my $disc (1 .. $D) {
        mkdir("d$disc",0);
        }

    chdir("..");
}
