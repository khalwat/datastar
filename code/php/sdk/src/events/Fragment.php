<?php
/**
 * @copyright Copyright (c) PutYourLightsOn
 */

namespace starfederation\datastar\events;

use starfederation\datastar\Defaults;
use starfederation\datastar\enums\EventType;
use starfederation\datastar\enums\FragmentMergeMode;

class Fragment implements EventInterface
{
    use EventTrait;

    public string $data;
    public ?string $selector = null;
    public ?FragmentMergeMode $mergeMode = null;
    public ?string $settleDuration = null;
    public ?bool $useViewTransition = null;

    public function __construct(string $data, array $options = [])
    {
        $this->data = $data;

        foreach ($options as $key => $value) {
            if (property_exists($this, $key)) {
                $this->$key = $value;
            }
        }
    }

    /**
     * @inerhitdoc
     */
    public function getEventType(): EventType
    {
        return EventType::Fragment;
    }

    /**
     * @inerhitdoc
     */
    public function getDataLines(): array
    {
        $data = trim($this->data);
        $dataLines = [];

        if ($this->selector !== null) {
            $dataLines[] = 'data: selector ' . $this->selector;
        }

        $dataLines[] = 'data: merge ' . ($this->mergeMode ?? Defaults::DEFAULT_FRAGMENT_MERGE_MODE->value);

        if ($this->settleDuration !== null && $this->settleDuration !== Defaults::DEFAULT_SETTLE_DURATION) {
            $dataLines[] = 'data: settleDuration ' . $this->settleDuration;
        }

        if ($this->useViewTransition === true) {
            $dataLines[] = 'data: useViewTransition true';
        }

        $dataLines[] = 'data: fragment';

        $lines = explode("\n", $data);
        foreach ($lines as $line) {
            $dataLines[] = 'data: ' . $line;
        }

        return $dataLines;
    }
}
