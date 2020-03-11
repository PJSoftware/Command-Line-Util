#!/perl -w

use strict;

print "Enter Season number: ";
my $S = <stdin>;
chomp $S;

my $E = 1;
my $Disc = 1;
my $ec;

mkdir("sx",0);
chdir("sx");

do {
    print "Enter episodes on disc $Disc (in S${S}E${E}): ";
    $ec = <stdin>;
    chomp $ec;
    if ($ec =~ /\d/) {
        print "Creating 'S${S}E$E'\n";
        mkdir("S${S}E$E",0);
        $E += $ec;
        $Disc++;
        }
    } while $ec;

chdir("..");
