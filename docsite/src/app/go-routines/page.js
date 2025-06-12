import SectionLayout from '../../components/SectionLayout';
import Introduction from './Introduction';
import Examples from './Examples';
import Exercises from './Exercises';
import Playground from './Playground';

export default function GoRoutines() {
    return (
        <SectionLayout
            title="Go Routines"
            introduction={<Introduction />}
            examples={<Examples />}
            exercises={<Exercises />}
            playground={<Playground />}
        />
    );
}